package child_chain

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"fmt"
	"github.com/renlulu/plasma-go/root-chain/artifact"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/mobile"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	artifact "github.com/renlulu/plasma-go/root-chain/artifact"
	"github.com/renlulu/plasma-go/core"
	"encoding/json"
)

const (
	Confirmations = 6
)

type RootChainListener struct {
	url          string
	rpcClient    *rpc.Client
	ethClient    *ethclient.Client
	RootChain    chain.RootChain
	chain        core.Chain
	handledEvent map[string]interface{}
}

func MakeRootChainListener(url string, rootChain chain.RootChain, ethUrl string, chain core.Chain) RootChainListener {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}

	ethClient, err := ethclient.Dial(ethUrl)

	if err != nil {
		log.Fatalf("could not create eth client: %v", err)
	}

	return RootChainListener{
		url:          url,
		rpcClient:    client,
		RootChain:    rootChain,
		ethClient:    ethClient,
		chain:        chain,
		handledEvent: make(map[string]interface{}, 0),
	}
}

func (listener *RootChainListener) GetLatestBlock() geth.Block {
	var latestBlock geth.Block
	err := listener.rpcClient.Call(&latestBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("can't get latest block:", err)
	}
	return latestBlock
}

func (listener *RootChainListener) EventListener(contract string) {
	latestBlock := listener.GetLatestBlock()
	contractAddress := common.HexToAddress(contract)
	var i = big.Int{}

	// 6 blocks ensure confirmation
	from := i.SetInt64(latestBlock.GetNumber() - (Confirmations*2 + 1))
	to := i.SetInt64(latestBlock.GetNumber() + 1 - Confirmations)

	q := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := listener.ethClient.SubscribeFilterLogs(context.Background(), q, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {

		case err := <-sub.Err():
			log.Fatal(err)

		case vLog := <-logs:
			contractAbi, err := abi.JSON(strings.NewReader(string(artifact.RootChainABI)))
			if err != nil {
				log.Fatal(err)
			}

			depositEvent := chain.RootChainDeposit{}
			exitStartedEvent := chain.RootChainExitStarted{}
			err = contractAbi.Unpack(&depositEvent, "RootChainDeposit", vLog.Data)
			if err != nil {
				log.Fatal(err)
				err = contractAbi.Unpack(&exitStartedEvent, "RootChainExitStarted", vLog.Data)
				if err != nil {
					log.Fatal(err)
				} else {
					utxoId := exitStartedEvent.UtxoPos
					listener.chain.MarkUTXOSpend(utxoId.Uint64())
					b, e := json.Marshal(exitStartedEvent)
					if e != nil {
						fmt.Printf("Error: %s", err)
						return
					}
					listener.handledEvent[string(b)] = exitStartedEvent
				}
			} else {
				//add to depositor block to child block
				owner := depositEvent.Depositor
				amount := depositEvent.Amount
				depositTx := core.MakeTransaction(owner, amount.Uint64())
				txs := make([]*core.Transaction, 0)
				txs = append(txs, depositTx)
				block := core.MakeBlock(txs, depositEvent.DepositBlock.Uint64())
				listener.chain.AddBlock(&block)
				b,e := json.Marshal(depositEvent)
				if e != nil {
					fmt.Printf("Error: %s", err)
					return
				}

				listener.handledEvent[string(b)] = depositEvent
			}

		}
	}
}

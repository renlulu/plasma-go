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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	artifact "github.com/renlulu/plasma-go/root-chain/artifact"
	"github.com/renlulu/plasma-go/child-chain/core"
	"encoding/json"
	"github.com/renlulu/plasma-go/child-chain/util"
	"strconv"
	"github.com/ethereum/go-ethereum/core/types"
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

func (listener *RootChainListener) GetLatestBlock() int64 {
	number := util.GetLatestBlock(listener.url).Result

	s := strings.Split(number,"x")
	_,number = s[0], s[1]
	latestNumber, err := strconv.ParseInt(number ,16, 64)
	if err != nil {
		fmt.Printf("convert error\n")
	}

	return latestNumber
}

func (listener *RootChainListener) EventListener(contract string) {
	latestBlock := listener.GetLatestBlock()
	fmt.Printf("event listener,latest block is %d",latestBlock)
	contractAddress := common.HexToAddress(contract)
	var i = big.Int{}

	// 6 blocks ensure confirmation
	from := i.SetInt64(latestBlock - (Confirmations*2 + 1))
	to := i.SetInt64(latestBlock + 1 - Confirmations)

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

package child

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/renlulu/plasma-go/child/core"
	"github.com/renlulu/plasma-go/child/util"
	"github.com/renlulu/plasma-go/root/artifact"
	artifact "github.com/renlulu/plasma-go/root/artifact"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
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
	fromBlock    int64
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
		fromBlock:    0,
	}
}

func (listener *RootChainListener) GetLatestBlock() int64 {
	number := util.GetLatestBlock(listener.url).Result

	s := strings.Split(number, "x")
	_, number = s[0], s[1]
	latestNumber, err := strconv.ParseInt(number, 16, 64)
	if err != nil {
		fmt.Printf("convert error\n")
	}

	return latestNumber
}

func (listener *RootChainListener) EventLoop(contract string) {
	for {
		fmt.Println("start event lintener...")
		go listener.EventListener(contract)

		time.Sleep(time.Second * 5)
	}
}

func (listener *RootChainListener) EventListener(contract string) {

	latestBlock := listener.GetLatestBlock()

	fmt.Println("start event listener, latest block is ", latestBlock)

	contractAddress := common.HexToAddress(contract)

	fmt.Println("contract address is ", contract)

	q := ethereum.FilterQuery{
		FromBlock: big.NewInt(listener.fromBlock),
		Addresses: []common.Address{contractAddress},
		ToBlock:   big.NewInt(latestBlock - Confirmations),
	}

	fmt.Printf("query is %+v\n", q)

	logs, err := listener.ethClient.FilterLogs(context.Background(), q)

	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(artifact.RootChainABI)))

	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		depositEvent := chain.RootChainDeposit{}
		exitStartedEvent := chain.RootChainExitStarted{}

		err = contractAbi.Unpack(&depositEvent, "Deposit", vLog.Data)
		if err != nil {
			log.Fatal(err)
			err = contractAbi.Unpack(&exitStartedEvent, "ExitStarted", vLog.Data)
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
			fmt.Printf("deposit block %+v\n", depositEvent)
			owner := depositEvent.Depositor
			amount := depositEvent.Amount
			depositTx := core.MakeTransaction(owner, amount.Uint64())
			txs := make([]*core.Transaction, 0)
			txs = append(txs, depositTx)
			block := core.MakeBlock(txs, depositEvent.DepositBlock.Uint64())
			listener.chain.AddBlock(&block)
			b, e := json.Marshal(depositEvent)
			if e != nil {
				fmt.Printf("Error: %s", err)
				return
			}

			listener.handledEvent[string(b)] = depositEvent
		}

		listener.fromBlock = listener.fromBlock + 1
	}
}

// It does't work, I don't know why
func (listener *RootChainListener) eventListener(contract string) {

	latestBlock := listener.GetLatestBlock()

	fmt.Println("start event listener, latest block is ", latestBlock)

	contractAddress := common.HexToAddress("0x44da3d92af236ffb5069781fa202c2d0e740d6a3")

	fmt.Println("contract address is ", contract)

	// 6 blocks ensure confirmation
	//from := i.SetInt64(latestBlock - (Confirmations * 2 + 1))
	//to := i.SetInt64(latestBlock + 1 - Confirmations)

	q := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := listener.ethClient.SubscribeFilterLogs(context.Background(), q, logs)

	if err != nil {
		log.Fatal(err)
	}

	for {

		fmt.Println("select...")

		select {

		case err := <-sub.Err():

			fmt.Println("get error ", err)
			log.Fatal(err)

		case vLog := <-logs:

			fmt.Println("get logs ")
			contractAbi, err := abi.JSON(strings.NewReader(string(artifact.RootChainABI)))
			if err != nil {
				log.Fatal(err)
			}

			depositEvent := chain.RootChainDeposit{}
			exitStartedEvent := chain.RootChainExitStarted{}
			err = contractAbi.Unpack(&depositEvent, "Deposit", vLog.Data)
			if err != nil {
				log.Fatal(err)
				err = contractAbi.Unpack(&exitStartedEvent, "ExitStarted", vLog.Data)
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
				//block := core.MakeBlock(txs, depositEvent.DepositBlock)
				//listener.chain.AddBlock(&block)
				//b, e := json.Marshal(depositEvent)
				//if e != nil {
				//	fmt.Printf("Error: %s", err)
				//	return
				//}
				//
				//listener.handledEvent[string(b)] = depositEvent
			}

		}
	}
}

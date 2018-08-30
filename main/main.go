package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/renlulu/plasma-go/child"
	"github.com/renlulu/plasma-go/cli"
	"github.com/renlulu/plasma-go/root/artifact"
	"log"
	"net/http"
)

var childChain child.ChildChain
var rootChain *chain.RootChain
var RootAddress = common.HexToAddress("0x44da3d92af236ffb5069781fa202c2d0e740d6a3")

func main() {
	var startServer bool
	var command string
	flag.BoolVar(&startServer, "server_start", false, "start http server?")
	flag.StringVar(&command, "run", "", "command")
	flag.Parse()

	if startServer {
		httpServer()
	}

	switch command {
	case "get_deposit_block":
		getDepositBlockC()
	}

}

func getDepositBlockC() {
	fmt.Println(cli.GetDepositBlock())
}

func httpServer() {
	http.HandleFunc("/get_current_block", child.GetCurrentBlockHandler)
	http.HandleFunc("/submit_block", child.SubmitBlockHandler)
	http.HandleFunc("/block_number", child.GetBlockNumberHandler)
	http.HandleFunc("/get_block", child.GetBlock)
	http.HandleFunc("/get_child_chain", child.GetChildChainHander)

	fmt.Printf("start http server on %s\n", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}

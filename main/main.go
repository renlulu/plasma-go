package main

import (
	"github.com/renlulu/plasma-go/child-chain"
	"net/http"
	"log"
	"fmt"
	"flag"
	"github.com/renlulu/plasma-go/cli"
)

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
	http.HandleFunc("/get_current_block", child_chain.GetCurrentBlockHandler)
	http.HandleFunc("/submit_block", child_chain.SubmitBlockHandler)
	http.HandleFunc("/block_number", child_chain.GetBlockNumberHandler)
	http.HandleFunc("/get_block", child_chain.GetBlock)
	http.HandleFunc("/get_child_chain", child_chain.GetChildChainHander)

	fmt.Printf("start http server on %s", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}

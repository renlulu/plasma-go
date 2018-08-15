package main

import (
	"net/http"
	"github.com/renlulu/plasma-go/child-chain"
	"fmt"
	"log"
)




func main() {
	http.HandleFunc("/get_current_block",child_chain.GetCurrentBlockHandler)
	http.HandleFunc("/submit_block",child_chain.SubmitBlockHandler)
	http.HandleFunc("/block_number",child_chain.GetBlockNumberHandler)
	http.HandleFunc("/get_block",child_chain.GetBlock)
	http.HandleFunc("/get_child_chain",child_chain.GetChildChainHander)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}
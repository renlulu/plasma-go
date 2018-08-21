package util

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var contentType = "application/json;charset=utf-8"


type Block struct {
	Number           string        `json:"number"`
	Hash             string        `json:"hash"`
	ParentHash       string        `json:"parentHash"`
	Nonce            string        `json:"nonce"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	LogsBloom        string        `json:"logsBloom"`
	TransactionsRoot string        `json:"transactionsRoot"`
	StateRoot        string        `json:"stateRoot"`
	Miner            string        `json:"miner"`
	Difficulty       string        `json:"difficulty""`
	TotalDifficulty  string        `json:"totalDifficulty"`
	ExtraData        string        `json:"extraData"`
	Size             string        `json:"size"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Timestamp        string        `json:"timestamp"`
	Transactions     []Transaction `json:"transactions"`
	Uncles           []string      `json:"uncles"`
}

type Transaction struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}


type BlockRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      byte          `json:"id"`
}

type BlockResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      byte          `json:"id"`
	Result  Block `json:"result"`
}

func GetLatestBlock(url string) BlockResponse {
	var block BlockResponse
	p := []interface{}{}
	request := BlockRequest{JsonRpc: "2.0", Method: "eth_blockNumber", Params: p, Id: 1}
	data, err := json.Marshal(request)

	fmt.Printf("request is %s\n",request)

	if err != nil {
		fmt.Println("json format error:", err)
		return block
	}

	body := bytes.NewBuffer(data)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		fmt.Println("Post failed:", err)
		return block
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return block
	}

	fmt.Println(string(content))

	json.Unmarshal(content, &block)

	return block
}

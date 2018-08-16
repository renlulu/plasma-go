package child_chain

import (
	"net/http"
	"encoding/json"
	"github.com/renlulu/plasma-go/core"
	"io/ioutil"
	"io"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	gcommon "github.com/go-ethereum/common"

	"github.com/renlulu/plasma-go/root-chain/artifact"
)


var childChain ChildChain
var rootChain *root_chain.RootChain
var CoinAddr = gcommon.HexToAddress("0x84F70FEa5Ba54323C0EF85c58A47c98E1a2fe2Db")



func init() {
	childChain = MakeChildChain(nil,nil)
	conn, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		panic(err)
	}
	rootChain,err = root_chain.NewRootChain(CoinAddr,conn)
	if err != nil {
		panic(err)
	}

}

type BlockNum struct {
	BlockNum uint64 `json:"block_number"`
}

func GetChildChainHander(w http.ResponseWriter, r *http.Request) {
	c := childChain.Chain
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(c)
}

func GetBlock(w http.ResponseWriter, r *http.Request) {
	var blockNum BlockNum

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &blockNum); err != nil {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	b := childChain.Chain.GetBlock(blockNum.BlockNum)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(b)
}

func GetCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	b := childChain.CurrentBlock
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(b)
}

func GetBlockNumberHandler(w http.ResponseWriter, r *http.Request) {
	n := childChain.GetBlockNum()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(n)

}

func SubmitBlockHandler(w http.ResponseWriter, r *http.Request) {
	var block core.Block
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &block); err != nil {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Printf("get new block %v, submit it\n",block)

	childChain.SubmitBlock(block)

}

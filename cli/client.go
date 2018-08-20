package cli

import (
	"github.com/renlulu/plasma-go/root-chain/artifact"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var rootChain *chain.RootChain
var RootAddress= common.HexToAddress("0xd917193e219ee3ced952c7fd4caae0b4a1b56567")
var connecter *Connecter

type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	rootchain *chain.RootChain
}

func init() {
	conn, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		panic(err)
	}
	rootChain,err = chain.NewRootChain(RootAddress,conn)
	if err != nil {
		panic(err)
	}

	connecter = &Connecter{
		ctx: context.Background(),
		conn: conn,
		rootchain:rootChain,
	}
}

func GetDepositBlock() uint64{
	opts := &bind.CallOpts{
		Pending:false,
		From: RootAddress,
		Context: connecter.ctx,
	}
	block,err := connecter.rootchain.GetDepositBlock(opts)
	if err != nil {
		panic(err)
	}
	return block.Uint64()
}
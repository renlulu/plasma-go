package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/renlulu/plasma-go/child-chain/core"
	"fmt"
)

type Provider struct {
	path  string
	db    *leveldb.DB
	chain *core.Chain
}

func NewProvider(path string) Provider {
	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		fmt.Printf("Init database failed")
	}

	return Provider{
		path: path, db: db,
	}
}

func (provider *Provider) Close() {
	err := provider.db.Close()
	if err != nil {
		fmt.Printf("Close database failed")
	}
}

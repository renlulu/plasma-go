package database

import (
	"fmt"
	"github.com/renlulu/plasma-go/child/core"
	"github.com/syndtr/goleveldb/leveldb"
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

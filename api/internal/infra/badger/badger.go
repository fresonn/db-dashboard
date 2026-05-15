package badgerdb

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

func New() *badger.DB {

	opts := badger.DefaultOptions("./data/badger")

	opts.MemTableSize = 128 << 20
	opts.BlockCacheSize = 256 << 20
	opts.IndexCacheSize = 128 << 20

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

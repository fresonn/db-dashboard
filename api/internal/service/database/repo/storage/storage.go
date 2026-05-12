package storage

import (
	"dashboard/api/internal/config"
	"dashboard/api/internal/infra/logger"
	"strconv"

	"github.com/dgraph-io/badger/v4"
)

type Storage struct {
	config config.AppConfig
	logger logger.Logger
	db     *badger.DB
}

func New(config config.AppConfig, logger logger.Logger, db *badger.DB) *Storage {
	return &Storage{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (s Storage) DatabaseName(id int) (string, error) {
	var result string

	err := s.db.View(func(txn *badger.Txn) error {
		key := []byte("db:" + strconv.Itoa(id) + ":name")

		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		value, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		result = string(value)

		return nil
	})

	return result, err
}

func (s Storage) SetDatabaseName(id int, name string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		key := []byte("db:" + strconv.Itoa(id) + ":name")
		value := []byte(name)

		return txn.Set(key, value)
	})
}

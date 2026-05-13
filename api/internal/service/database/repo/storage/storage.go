package storage

import (
	"dashboard/api/internal/config"
	"dashboard/api/internal/infra/logger"

	"github.com/dgraph-io/badger/v4"
)

// db:{id}:stat:overview:{timestamp}
// db:{id}:metric:size:10s:{timestamp}

// func paddedTimestamp(ts int64) string {
// 	return fmt.Sprintf("%020d", ts)
// }

// func MetricKey(
// 	dbID int64,
// 	metric string,
// 	resolution string,
// 	ts int64,
// ) []byte {
// 	return []byte(fmt.Sprintf(
// 		"db:%d:metric:%s:%s:%s",
// 		dbID,
// 		metric,
// 		resolution,
// 		paddedTimestamp(ts),
// 	))
// }

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

package storage

import (
	"dashboard/api/internal/model/database"
	service "dashboard/api/internal/service/database"
	pb "dashboard/api/internal/service/database/repo/storage/codec"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"google.golang.org/protobuf/proto"
)

/*
STAT:
db:47020:stat:overview        → current value

METRIC:
db:47020:metric:size:{ts}     → history
*/

func overviewStatsKey(id int) []byte {
	return fmt.Appendf([]byte{}, "db:%d:stat:overview", id)
}

func (s *Storage) SaveStatsOverview(databaseID int, stats database.PostgresDbOverviewStats) error {

	key := overviewStatsKey(databaseID)

	pb := EncodeOverviewStats(stats)

	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	fmt.Println("key", string(key))

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})
}

func (s *Storage) StatsOverview(databaseID int, stats database.PostgresDbOverviewStats) (database.PostgresDbOverviewStats, error) {

	key := overviewStatsKey(databaseID)

	var result database.PostgresDbOverviewStats

	err := s.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {

			var pbData pb.PostgresDbOverviewStats

			if err := proto.Unmarshal(val, &pbData); err != nil {
				return err
			}

			result = DecodeOverviewStats(&pbData)
			return nil
		})
	})

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return database.PostgresDbOverviewStats{}, service.ErrStatsOverviewNotFound
		}
		return database.PostgresDbOverviewStats{}, err
	}

	return result, nil
}

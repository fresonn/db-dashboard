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

func (s *Storage) SaveStatsOverview(databaseID int, stats database.StoredOverviewStats) error {

	key := overviewStatsKey(databaseID)

	pb := EncodeStoredOverviewStats(stats)

	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})
}

func (s *Storage) StatsOverview(databaseID int) (database.StoredOverviewStats, error) {

	key := overviewStatsKey(databaseID)

	var result database.StoredOverviewStats

	err := s.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {

			var pbData pb.StoredOverviewStats

			if err := proto.Unmarshal(val, &pbData); err != nil {
				return err
			}

			result = DecodeStoredOverviewStats(&pbData)
			return nil
		})
	})

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return database.StoredOverviewStats{}, service.ErrStatsOverviewNotFound
		}
		return database.StoredOverviewStats{}, err
	}

	return result, nil
}

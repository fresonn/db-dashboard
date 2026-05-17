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
db:47020:stat:{domain}        → current value

METRIC:
db:47020:metric:size:{ts}     → history
*/

func overviewStatKey(databaseID int, metric database.Metric) []byte {
	return fmt.Appendf([]byte{}, "db:%d:stat:%s", databaseID, metric)
}

func (s *Storage) SaveOverviewStatState(databaseID int, metric database.Metric, stat database.OverviewStatState) error {

	key := overviewStatKey(databaseID, metric)

	pb := EncodeOverviewStatState(stat)

	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})
}

func (s *Storage) OverviewStatState(databaseID int, metric database.Metric) (database.OverviewStatState, error) {

	key := overviewStatKey(databaseID, metric)

	var result database.OverviewStatState

	err := s.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {

			var pbData pb.OverviewStatState

			if err := proto.Unmarshal(val, &pbData); err != nil {
				return err
			}

			result = DecodeOverviewStatState(&pbData)
			return nil
		})
	})

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return database.OverviewStatState{}, service.ErrOverviewStatStateNotFound
		}
		return database.OverviewStatState{}, err
	}

	return result, nil
}

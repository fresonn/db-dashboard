package storage

import (
	"dashboard/api/internal/model/database"
	"time"

	pb "dashboard/api/internal/service/database/repo/storage/codec"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EncodeOverviewStats(
	m database.PostgresDbOverviewStats,
) *pb.PostgresDbOverviewStats {

	return &pb.PostgresDbOverviewStats{
		Id:      m.ID,
		Name:    m.Name,
		Size:    m.Size,
		Tables:  m.Tables,
		Indexes: m.Indexes,
	}
}

func DecodeOverviewStats(pbMsg *pb.PostgresDbOverviewStats) database.PostgresDbOverviewStats {

	if pbMsg == nil {
		return database.PostgresDbOverviewStats{}
	}

	return database.PostgresDbOverviewStats{
		ID:      pbMsg.Id,
		Name:    pbMsg.Name,
		Size:    pbMsg.Size,
		Tables:  pbMsg.Tables,
		Indexes: pbMsg.Indexes,
	}
}

func EncodeStoredOverviewStats(m database.StoredOverviewStats) *pb.StoredOverviewStats {

	return &pb.StoredOverviewStats{
		Current:   EncodeOverviewStats(m.Current),
		Previous:  EncodeOverviewStats(m.Previous),
		ChangedAt: timestamppb.New(m.ChangedAt),
	}
}

func DecodeStoredOverviewStats(pbMsg *pb.StoredOverviewStats) database.StoredOverviewStats {

	if pbMsg == nil {
		return database.StoredOverviewStats{}
	}

	var changedAt time.Time

	if pbMsg.ChangedAt != nil {
		changedAt = pbMsg.ChangedAt.AsTime()
	}

	return database.StoredOverviewStats{
		Current:   DecodeOverviewStats(pbMsg.Current),
		Previous:  DecodeOverviewStats(pbMsg.Previous),
		ChangedAt: changedAt,
	}
}

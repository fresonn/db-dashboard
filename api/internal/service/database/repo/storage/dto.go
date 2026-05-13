package storage

import (
	"dashboard/api/internal/model/database"

	pb "dashboard/api/internal/service/database/repo/storage/codec"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func EncodeOverviewStats(m database.PostgresDbOverviewStats) *pb.PostgresDbOverviewStats {
	return &pb.PostgresDbOverviewStats{
		Id:        m.ID,
		Name:      m.Name,
		Size:      m.Size,
		Tables:    m.Tables,
		Indexes:   m.Indexes,
		CreatedAt: timestamppb.Now(),
	}
}

func DecodeOverviewStats(pbMsg *pb.PostgresDbOverviewStats) database.PostgresDbOverviewStats {
	return database.PostgresDbOverviewStats{
		ID:      pbMsg.Id,
		Name:    pbMsg.Name,
		Size:    pbMsg.Size,
		Tables:  pbMsg.Tables,
		Indexes: pbMsg.Indexes,
	}
}

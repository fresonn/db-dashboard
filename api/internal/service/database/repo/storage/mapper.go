package storage

import (
	"dashboard/api/internal/model/database"
	"time"

	pb "dashboard/api/internal/service/database/repo/storage/codec"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EncodeOverviewStatState(m database.OverviewStatState) *pb.OverviewStatState {

	return &pb.OverviewStatState{
		Current:   m.Current,
		Previous:  m.Previous,
		ChangedAt: timestamppb.New(m.ChangedAt),
	}
}

func DecodeOverviewStatState(pbMsg *pb.OverviewStatState) database.OverviewStatState {

	if pbMsg == nil {
		return database.OverviewStatState{}
	}

	var createdAt time.Time

	if pbMsg.ChangedAt != nil {
		createdAt = pbMsg.ChangedAt.AsTime()
	}

	return database.OverviewStatState{
		Current:   pbMsg.Current,
		Previous:  pbMsg.Previous,
		ChangedAt: createdAt,
	}
}

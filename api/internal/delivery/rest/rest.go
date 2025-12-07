package rest

import (
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/scopes/cluster"
)

type Handler struct {
	cluster *cluster.Cluster
}

var _ openapi.StrictServerInterface = (*Handler)(nil)

func New(cluster *cluster.Cluster) *Handler {
	return &Handler{
		cluster: cluster,
	}
}

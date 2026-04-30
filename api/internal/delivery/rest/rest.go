package rest

import (
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/service/cluster"
	"dashboard/api/internal/service/database"
	"dashboard/api/internal/service/roles"
)

type Handler struct {
	cluster  cluster.Service
	roles    roles.Service
	database database.Service
}

var _ openapi.StrictServerInterface = (*Handler)(nil)

func New(cluster cluster.Service, roles roles.Service, database database.Service) *Handler {
	return &Handler{
		cluster:  cluster,
		roles:    roles,
		database: database,
	}
}

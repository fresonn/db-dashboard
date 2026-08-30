package http

import (
	"dashboard/api/gen/openapi"
)

type Handler struct {
	useCase         UseCases
	clusterService  ClusterService
	roleService     RoleService
	databaseService DatabaseService
}

var _ openapi.StrictServerInterface = (*Handler)(nil)

func New(useCase UseCases, clusterService ClusterService, roleService RoleService, databaseService DatabaseService) *Handler {
	return &Handler{
		useCase:         useCase,
		clusterService:  clusterService,
		roleService:     roleService,
		databaseService: databaseService,
	}
}

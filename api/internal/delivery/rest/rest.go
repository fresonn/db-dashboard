package rest

import (
	"context"
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/scopes/cluster"
	"dashboard/api/internal/scopes/cluster/entities"
	"errors"

	"github.com/go-playground/validator/v10"
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

func (h *Handler) GetStatus(ctx context.Context, request openapi.GetStatusRequestObject) (openapi.GetStatusResponseObject, error) {

	status := h.cluster.PostgresStatus(ctx)

	resp := openapi.GetStatusResponse{
		PostgresConnection: openapi.ConnectionStatus(status),
	}

	return openapi.GetStatus200JSONResponse(resp), nil
}

func (h *Handler) ClusterConnect(ctx context.Context, request openapi.ClusterConnectRequestObject) (openapi.ClusterConnectResponseObject, error) {

	err := h.cluster.Connect(ctx, entities.AuthData(*request.Body))

	if err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return openapi.ClusterConnect422JSONResponse{
				Message: err.Error(),
			}, nil
		}

		return openapi.ClusterConnect400JSONResponse{
			Error: err.Error(),
		}, nil
	}

	resp := openapi.GetStatusResponse{
		PostgresConnection: openapi.PgConnectionStatusActive,
	}

	return openapi.ClusterConnect200JSONResponse(resp), nil
}

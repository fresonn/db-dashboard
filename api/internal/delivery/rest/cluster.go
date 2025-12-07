package rest

import (
	"context"
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/scopes/cluster/entities"
	"errors"

	"github.com/go-playground/validator/v10"
)

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
				Message: "Request validation failed",
				Reason:  err.Error(),
			}, nil
		}

		return openapi.ClusterConnect400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.ClusterConnect200JSONResponse{
		PostgresConnection: openapi.PgConnectionStatusConnected,
	}, nil
}

func (h *Handler) ClusterDisconnect(ctx context.Context, _request openapi.ClusterDisconnectRequestObject) (openapi.ClusterDisconnectResponseObject, error) {

	err := h.cluster.Disconnect(ctx)
	if err != nil {
		return openapi.ClusterDisconnect400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.ClusterDisconnect200JSONResponse{
		PostgresConnection: openapi.PgConnectionStatusDisconnected,
	}, nil
}

func (h *Handler) PostgresVersion(ctx context.Context, request openapi.PostgresVersionRequestObject) (openapi.PostgresVersionResponseObject, error) {

	version, err := h.cluster.Version(ctx)
	if err != nil {
		return openapi.PostgresVersion400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.PostgresVersion200JSONResponse(version), nil
}

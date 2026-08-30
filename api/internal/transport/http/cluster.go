package http

import (
	"context"
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/model/cluster"
	"errors"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) GetStatus(ctx context.Context, req openapi.GetStatusRequestObject) (openapi.GetStatusResponseObject, error) {

	status := h.clusterService.PostgresStatus(ctx)

	resp := openapi.GetStatusResponse{
		User:             status.CurrentUser,
		Database:         status.CurrentDatabase,
		ConnectionStatus: openapi.ConnectionStatus(status.ConnectionStatus),
	}

	return openapi.GetStatus200JSONResponse(resp), nil
}

func (h *Handler) SwitchCurrentDatabase(
	ctx context.Context,
	req openapi.SwitchCurrentDatabaseRequestObject,
) (openapi.SwitchCurrentDatabaseResponseObject, error) {

	status, err := h.useCase.SwitchDatabase(ctx, req.Body.DatabaseId)
	if err != nil {
		return openapi.SwitchCurrentDatabase400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	resp := openapi.GetStatusResponse{
		User:             status.CurrentUser,
		Database:         status.CurrentDatabase,
		ConnectionStatus: openapi.ConnectionStatus(status.ConnectionStatus),
	}

	return openapi.SwitchCurrentDatabase200JSONResponse(resp), nil
}

func (h *Handler) ClusterConnect(ctx context.Context, req openapi.ClusterConnectRequestObject) (openapi.ClusterConnectResponseObject, error) {

	status, err := h.clusterService.Connect(ctx, cluster.NewConnection(*req.Body))
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
		User:             status.CurrentUser,
		Database:         status.CurrentDatabase,
		ConnectionStatus: openapi.PgConnectionStatusConnected,
	}, nil
}

func (h *Handler) ClusterDisconnect(ctx context.Context, req openapi.ClusterDisconnectRequestObject) (openapi.ClusterDisconnectResponseObject, error) {

	err := h.clusterService.Disconnect(ctx)
	if err != nil {
		return openapi.ClusterDisconnect400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.ClusterDisconnect200JSONResponse{
		ConnectionStatus: openapi.PgConnectionStatusDisconnected,
	}, nil
}

func (h *Handler) PostgresVersion(ctx context.Context, req openapi.PostgresVersionRequestObject) (openapi.PostgresVersionResponseObject, error) {

	version, err := h.clusterService.Version(ctx)
	if err != nil {
		return openapi.PostgresVersion400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.PostgresVersion200JSONResponse(version), nil
}

func (h *Handler) PostgresUptime(ctx context.Context, req openapi.PostgresUptimeRequestObject) (openapi.PostgresUptimeResponseObject, error) {

	uptime, err := h.clusterService.Uptime(ctx)
	if err != nil {
		return openapi.PostgresUptime400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.PostgresUptime200JSONResponse(uptime), nil
}

func (h *Handler) PostmasterSettings(ctx context.Context, req openapi.PostmasterSettingsRequestObject) (openapi.PostmasterSettingsResponseObject, error) {

	settings, err := h.clusterService.PostmasterSettings(ctx)
	if err != nil {
		return openapi.PostmasterSettings400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.PostmasterSettings200JSONResponse(settings), nil
}

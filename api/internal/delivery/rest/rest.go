package rest

import (
	"context"
	"dashboard/api/gen/openapi"
)

type Handler struct{}

var _ openapi.StrictServerInterface = (*Handler)(nil)

func New() *Handler {
	return &Handler{}
}

func (h *Handler) GetStatus(ctx context.Context, request openapi.GetStatusRequestObject) (openapi.GetStatusResponseObject, error) {

	resp := openapi.GetStatusResponse{
		PostgresConnection: openapi.PgConnectionStatusConnecting,
	}

	return openapi.GetStatus200JSONResponse(resp), nil
}

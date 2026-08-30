package http

import (
	"context"
	"dashboard/api/gen/openapi"
)

func (h *Handler) Roles(ctx context.Context, req openapi.RolesRequestObject) (openapi.RolesResponseObject, error) {

	roles, err := h.roleService.Roles(ctx)
	if err != nil {
		return openapi.Roles400JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return openapi.Roles200JSONResponse(roles), nil
}

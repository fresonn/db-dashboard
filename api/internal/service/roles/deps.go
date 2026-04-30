package roles

import (
	"context"
	"dashboard/api/internal/model/role"
)

type Storage interface {
	Roles(ctx context.Context) ([]role.Role, error)
}

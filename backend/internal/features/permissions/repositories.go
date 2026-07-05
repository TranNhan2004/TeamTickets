package permissions

import (
	"context"

	"github.com/google/uuid"
)

type PermissionRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Permission, error)
	FindByCode(ctx context.Context, code string) (*Permission, error)
}

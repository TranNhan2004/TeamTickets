package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
)

type UserRepository interface {
	Create(ctx context.Context, user *User, audit *shared.AuditModel) error
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, user *User, audit *shared.AuditModel) error
	Delete(ctx context.Context, id uuid.UUID, audit *shared.AuditModel) error
}

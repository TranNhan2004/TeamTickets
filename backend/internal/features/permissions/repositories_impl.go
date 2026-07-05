package permissions

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
	"gorm.io/gorm"
)

type permissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepositoryImpl{
		db: db,
	}
}

func (p *permissionRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*Permission, error) {
	db := shared.DBFromContext(ctx, p.db)

	var permission Permission
	err := db.
		WithContext(ctx).
		Where("id = ? AND is_deleted = false", id).
		First(&permission).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (p *permissionRepositoryImpl) FindByCode(ctx context.Context, code string) (*Permission, error) {
	db := shared.DBFromContext(ctx, p.db)

	var permission Permission
	err := db.
		WithContext(ctx).
		Where("code = ? AND is_deleted = false", code).
		First(&permission).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

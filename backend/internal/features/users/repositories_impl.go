package users

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *User, audit *shared.AuditModel) error {
	db := shared.DBFromContext(ctx, r.db)

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user.ID = id
	user.CreatedAt = audit.AuditTime
	user.UpdatedAt = audit.AuditTime

	return db.
		WithContext(ctx).
		Create(user).
		Error
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	db := shared.DBFromContext(ctx, r.db)

	var user User
	err := db.
		WithContext(ctx).
		Where("id = ? AND is_deleted = false", id).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *User, audit *shared.AuditModel) error {
	db := shared.DBFromContext(ctx, r.db)

	user.UpdatedAt = audit.AuditTime
	result := db.
		WithContext(ctx).
		Model(&User{}).
		Where("id = ? AND is_deleted = false", user.ID).
		Updates(map[string]any{
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"display_name":      user.DisplayName,
			"email":             user.Email,
			"hashed_password":   user.HashedPassword,
			"is_email_verified": user.IsEmailVerified,
			"is_active":         user.IsActive,
			"avatar_url":        user.AvatarURL,
			"updated_at":        user.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id uuid.UUID, audit *shared.AuditModel) error {
	db := shared.DBFromContext(ctx, r.db)

	result := db.
		WithContext(ctx).
		Model(&User{}).
		Where("id = ? AND is_deleted = false", id).
		Updates(map[string]any{
			"is_deleted": true,
			"deleted_at": audit.AuditTime,
			"updated_at": audit.AuditTime,
			"is_active":  false,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

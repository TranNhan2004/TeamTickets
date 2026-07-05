package users

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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

func (r *userRepositoryImpl) Create(ctx context.Context, user *User) error {
	if user.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}

		user.ID = id
	}

	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.
		WithContext(ctx).
		Create(user).
		Error
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var user User

	err := r.db.
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

func (r *userRepositoryImpl) Update(ctx context.Context, user *User) error {
	user.UpdatedAt = time.Now().UTC()

	result := r.db.
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

func (r *userRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	now := time.Now().UTC()

	result := r.db.
		WithContext(ctx).
		Model(&User{}).
		Where("id = ? AND is_deleted = false", id).
		Updates(map[string]any{
			"is_deleted": true,
			"deleted_at": now,
			"updated_at": now,
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

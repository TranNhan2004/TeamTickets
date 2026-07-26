package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	IdentitySubject string    `gorm:"column:identity_subject"`
	FirstName       string    `gorm:"column:first_name"`
	LastName        string    `gorm:"column:last_name"`
	DisplayName     string    `gorm:"column:display_name"`
	Email           string    `gorm:"column:email"`

	// Deprecated: HashedPassword is deprecated and will be removed in future versions
	HashedPassword string `gorm:"column:hashed_password"`

	AvatarURL       *string    `gorm:"column:avatar_url"`
	IsEmailVerified bool       `gorm:"column:is_email_verified"`
	IsActive        bool       `gorm:"column:is_active"`
	LastLoginAt     *time.Time `gorm:"column:last_login_at"`
	LastLogoutAt    *time.Time `gorm:"column:last_logout_at"`
	IsDeleted       bool       `gorm:"column:is_deleted"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at"`
}

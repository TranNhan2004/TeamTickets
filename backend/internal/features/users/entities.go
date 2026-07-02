package users

import "time"

type User struct {
	ID              string
	FirstName       string
	LastName        string
	DisplayName     string
	Email           string
	HashedPassword  string
	IsEmailVerified bool
	IsActive        bool
	AvatarURL       *string
	DeletedAt       *time.Time
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

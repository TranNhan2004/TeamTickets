package users

import (
	"time"
)

// Users
type CreateUserRequest struct {
	FirstName       string  `json:"firstName" validate:"required"`
	LastName        string  `json:"lastName" validate:"required"`
	DisplayName     string  `json:"displayName"`
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required,min=8"`
	IsEmailVerified bool    `json:"isEmailVerified"`
	IsActive        bool    `json:"isActive"`
	AvatarURL       *string `json:"avatarURL"`
}

func (r *CreateUserRequest) ToModel() *CreateUserRequestModel {
	return &CreateUserRequestModel{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		DisplayName:     r.DisplayName,
		Email:           r.Email,
		Password:        r.Password,
		IsEmailVerified: r.IsEmailVerified,
		IsActive:        r.IsActive,
		AvatarURL:       r.AvatarURL,
	}
}

type CreateUserResponse struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	AvatarURL   *string   `json:"avatarURL"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateUserRequest struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	DisplayName string  `json:"displayName"`
	AvatarURL   *string `json:"avatarURL"`
}

func (r *UpdateUserRequest) ToModel() *UpdateUserRequestModel {
	return &UpdateUserRequestModel{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		DisplayName: r.DisplayName,
		AvatarURL:   r.AvatarURL,
	}
}

type UpdateUserResponse struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	AvatarURL   *string   `json:"avatarURL"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetUserResponse struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	AvatarURL   *string   `json:"avatarURL"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

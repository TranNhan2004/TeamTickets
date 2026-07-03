package users

import (
	"time"

	"github.com/google/uuid"
)

// Users
type CreateUserRequestModel struct {
	FirstName       string
	LastName        string
	DisplayName     string
	Email           string
	Password        string
	IsEmailVerified bool
	IsActive        bool
	AvatarURL       *string
}

type CreateUserResponseModel struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DisplayName string
	Email       string
	AvatarURL   *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *CreateUserResponseModel) ToDTO() *CreateUserResponse {
	return &CreateUserResponse{
		ID:          r.ID.String(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		DisplayName: r.DisplayName,
		Email:       r.Email,
		AvatarURL:   r.AvatarURL,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

type UpdateUserRequestModel struct {
	FirstName   string
	LastName    string
	DisplayName string
	AvatarURL   *string
}

type UpdateUserResponseModel struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DisplayName string
	Email       string
	AvatarURL   *string
	UpdatedAt   time.Time
}

func (r *UpdateUserResponseModel) ToDTO() *UpdateUserResponse {
	return &UpdateUserResponse{
		ID:          r.ID.String(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		DisplayName: r.DisplayName,
		Email:       r.Email,
		AvatarURL:   r.AvatarURL,
		UpdatedAt:   r.UpdatedAt,
	}
}

type GetUserResponseModel struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DisplayName string
	Email       string
	AvatarURL   *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *GetUserResponseModel) ToDTO() *GetUserResponse {
	return &GetUserResponse{
		ID:          r.ID.String(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		DisplayName: r.DisplayName,
		Email:       r.Email,
		AvatarURL:   r.AvatarURL,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

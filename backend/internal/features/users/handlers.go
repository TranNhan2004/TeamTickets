package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/httphelpers"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var request CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	user, err := h.userService.Create(ctx.Request.Context(), CreateUserRequestModel{
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		DisplayName:     request.DisplayName,
		Email:           request.Email,
		Password:        request.Password,
		IsEmailVerified: request.IsEmailVerified,
		IsActive:        request.IsActive,
		AvatarURL:       request.AvatarURL,
	})

	if err != nil {
		apperrors.Respond(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, user.ToDTO())
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	id, err := httphelpers.ParseUUIDParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	user, err := h.userService.FindByID(ctx.Request.Context(), id)
	if err != nil {
		apperrors.Respond(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToDTO())
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := httphelpers.ParseUUIDParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	var request UpdateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	user, err := h.userService.Update(ctx.Request.Context(), id, UpdateUserRequestModel{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DisplayName: request.DisplayName,
		AvatarURL:   request.AvatarURL,
	})

	if err != nil {
		apperrors.Respond(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToDTO())
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, err := httphelpers.ParseUUIDParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	if err := h.userService.Delete(ctx.Request.Context(), id); err != nil {
		apperrors.Respond(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

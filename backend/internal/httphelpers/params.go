package httphelpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
)

func ParseUUIDParam(ctx *gin.Context, paramName string) (uuid.UUID, *apperrors.AppError) {
	rawID := ctx.Param(paramName)

	id, err := uuid.Parse(rawID)
	if err != nil {
		return uuid.Nil, apperrors.BadRequest("INVALID_UUID", "invalid UUID format")
	}

	return id, nil
}

package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}

func Respond(ctx *gin.Context, err error) {
	var appErr *AppError

	if errors.As(err, &appErr) {
		ctx.JSON(appErr.StatusCode, ErrorResponse{
			StatusCode: appErr.StatusCode,
			Code:       appErr.Code,
			Message:    appErr.Message,
			Details:    appErr.Details,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
	})
}

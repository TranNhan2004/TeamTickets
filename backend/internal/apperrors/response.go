package apperrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func Respond(ctx *gin.Context, appErr *AppError) {
	if appErr != nil {
		ctx.JSON(appErr.StatusCode, ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Details: appErr.Details,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Something went wrong",
	})
}

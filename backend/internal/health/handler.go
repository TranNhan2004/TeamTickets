package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Live(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) Ready(ctx *gin.Context) {
	checkCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second)
	defer cancel()

	results, healthy := h.service.CheckReadiness(checkCtx)

	if !healthy {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"checks": results,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"checks": results,
	})
}

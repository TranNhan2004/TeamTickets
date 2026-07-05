package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
)

func MustGetCurrentlyUser(c *gin.Context) (*shared.CurrentUserModel, bool) {
	currentUser, ok := shared.GetCurrentUser(c)
	if ok {
		return currentUser, true
	}

	unauthorizedErr := NewErrorUnauthorized()
	apperrors.RespondAbort(c, unauthorizedErr)

	return nil, false
}

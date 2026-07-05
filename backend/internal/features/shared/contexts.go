package shared

import (
	"github.com/gin-gonic/gin"
)

const GinCurrentUserKey = "current_user"

func SetCurrentUser(c *gin.Context, u *CurrentUserModel) {
	c.Set(GinCurrentUserKey, u)
}

func GetCurrentUser(c *gin.Context) (*CurrentUserModel, bool) {
	value, exists := c.Get(GinCurrentUserKey)
	if !exists {
		return nil, false
	}

	authUser, ok := value.(*CurrentUserModel)
	if !ok || authUser == nil {
		return nil, false
	}

	return authUser, true
}

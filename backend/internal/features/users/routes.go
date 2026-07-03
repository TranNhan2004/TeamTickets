package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	users := rg.Group("/users")
	{
		users.POST("", handler.Create)
		users.GET("/:id", handler.FindByID)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
	}
}

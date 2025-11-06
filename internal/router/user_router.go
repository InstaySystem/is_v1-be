package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup, hdl *handler.UserHandler, authMid *middleware.AuthMiddleware) {
	user := rg.Group("/users", authMid.IsAuthentication())
	{
		user.POST("", authMid.HasAnyRole([]string{"admin"}), hdl.CreateUser)

		user.GET("/:id", authMid.HasAnyRole([]string{"admin"}), hdl.GetUserByID)

		user.GET("", authMid.HasAnyRole([]string{"admin"}), hdl.GetUsers)

		user.GET("/roles", authMid.HasAnyRole([]string{"admin"}), hdl.GetAllRoles)

		user.PATCH("/:id", authMid.HasAnyRole([]string{"admin"}), hdl.UpdateUser)

		user.PUT("/:id/password", authMid.HasAnyRole([]string{"admin"}), hdl.UpdateUserPassword)

		user.DELETE("/:id", authMid.HasAnyRole([]string{"admin"}), hdl.DeleteUser)
	}
}

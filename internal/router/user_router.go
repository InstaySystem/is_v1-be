package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup, hdl *handler.UserHandler) {
	user := rg.Group("/users")
	{
		user.POST("", hdl.CreateUser)
	}
}
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolist/api"
)

func NewRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())
	}
	return r
}
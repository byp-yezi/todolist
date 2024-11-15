package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolist/api"
	"todolist/middleware"
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
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			auth.POST("task_create", api.CreateTaskHandler())
			auth.POST("task_list", api.ListTaskHandler())
			auth.POST("task_show", api.ShowTaskHandler())
		}
	}
	return r
}

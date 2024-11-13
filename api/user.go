package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolist/pkg/util"
	"todolist/service"
	"todolist/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			l := service.GetUserSrv()
			resp, err := l.Register(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}

	}
}
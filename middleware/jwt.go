package middleware

import (
	"net/http"
	"time"
	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusForbidden
			ctx.JSON(e.InvalidParams, ctl.RespError(nil, "缺少token", code))
			ctx.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthTokenFail
		} else if claims.ExpiresAt < time.Now().Unix() {
			code = e.ErrorAuthTokenExpired
		}
		if code != e.SUCCESS {
			ctx.JSON(e.InvalidParams, ctl.RespError(nil, "可能是身份验证过期，请重新登录", code))
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		ctx.Next()
	}
}
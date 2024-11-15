package middleware

import (
	"todolist/pkg/e"
	"todolist/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			code = e.ErrorAuthNotFound
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg": e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthTokenFail
			util.LogrusObj.Infoln(err.Error())
		}
		if code != e.SUCCESS {
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg": e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		// ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		ctx.Set("userId", claims.Id)
		ctx.Next()
	}
}
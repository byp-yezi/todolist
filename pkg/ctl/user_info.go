package ctl

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// type key int

// var userKey key

// type UserInfo struct {
// 	Id uint `json:"id"`
// }

// func NewContext(ctx context.Context, u *UserInfo) context.Context {
// 	return context.WithValue(ctx, userKey, u)
// }

func GetUserId(ctx *gin.Context) (uint, error) {
	userIdRaw, ok := ctx.Get("userId")
	if !ok {
		return 0, errors.New("获取用户信息错误")
	}
	userId, ok := userIdRaw.(uint)
	if !ok {
		return 0, errors.New("用户信息类型错误")
	}
	return userId, nil
}
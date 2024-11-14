package ctl

import "context"

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}
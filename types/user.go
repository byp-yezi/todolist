package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"yezi"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"yezi666"`
}

type UserServiceResp struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	CreateAt int64  `json:"create_at" form:"create_at"`
	Token    string `json:"token" form:"token"`
}

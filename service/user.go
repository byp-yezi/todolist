package service

import (
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"

	"todolist/pkg/ctl"
	"todolist/pkg/util"
	"todolist/repository/db/dao"
	"todolist/repository/db/model"
	"todolist/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct{}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (userSrv *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, err = userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		u := &model.User{
			UserName: req.UserName,
		}
		if err = u.SetPassword(req.Password); err != nil {
			util.LogrusObj.Info(err)
			return
		}
		if err = userDao.CreateUser(u); err != nil {
			util.LogrusObj.Info(err)
			return
		}
		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("the user is exist")
		return
	default:
		return
	}
}
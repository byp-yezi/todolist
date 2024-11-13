package dao

import (
	"context"
	"todolist/repository/db/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (userDao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = userDao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	return
}

func (userDao *UserDao) CreateUser(user *model.User) (err error) {
	err = userDao.DB.Model(&model.User{}).Create(user).Error
	return
}

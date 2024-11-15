package dao

import (
	"context"

	"gorm.io/gorm"

	"todolist/repository/db/model"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (taskDao *TaskDao) CreateTask(task *model.Task) error {
	return taskDao.Model(&model.Task{}).Create(&task).Error
}

func (taskDao *TaskDao) ListTask(page, pageSize int, userId uint) (tasks []*model.Task, total int64, err error) {
	err = taskDao.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&tasks).Error
	return
}

func (taskDao *TaskDao) ShowTask(id, userId uint) (task *model.Task, err error) {
	err = taskDao.Model(&model.Task{}).Where("id = ? AND uid = ?", id, userId).First(&task).Error
	return
}

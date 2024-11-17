package dao

import (
	"context"
	"errors"

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
	rows := taskDao.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&tasks).RowsAffected
	if rows == 0 {
		err = errors.New("任务列表为空")
		return
	}
	return
}

func (taskDao *TaskDao) FindTaskByIdAndUserId(id, userId uint) (task *model.Task, err error) {
	err = taskDao.Model(&model.Task{}).Where("id = ? AND uid = ?", id, userId).First(&task).Error
	return
}

func (taskDao *TaskDao) SearchTask(info string, userId uint) (tasks []*model.Task, err error) {
	rows := taskDao.Model(&model.Task{}).Where("uid = ? AND (title like ? OR content like ?)", userId, "%"+info+"%", "%"+info+"%").Find(&tasks).RowsAffected
	if rows == 0 {
		err = errors.New("搜索结果为空")
		return
	}
	return
}

func (taskDao *TaskDao) DeleteTask(id, userId uint) error {
	task, err := taskDao.FindTaskByIdAndUserId(id ,userId)
	if err != nil {
		return err
	}
	return taskDao.Delete(&task).Error
}

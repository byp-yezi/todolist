package model

import (
	"strconv"
	"todolist/repository/cache"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}

// 获取点击数
func (task *Task) GetView() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(task.ID)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return uint64(count)
}

func (task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(task.ID))
}
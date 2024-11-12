package dao

import "todolist/repository/db/model"

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
}
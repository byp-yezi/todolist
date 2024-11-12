package main

import (
	conf "todolist/config"
	"todolist/pkg/util"
	"todolist/repository/cache"
	"todolist/repository/db/dao"
)

func main() {
	loading()
}

func loading() {
	conf.Init()
	util.InitLog()
	dao.MysqlInit()
	cache.RedisInit()
}
package main

import (
	conf "todolist/config"
	"todolist/pkg/util"
	"todolist/repository/cache"
	"todolist/repository/db/dao"
	"todolist/routes"
)

func main() {
	loading()
	r := routes.NewRoutes()
	_ = r.Run(conf.HttpPort)
}

func loading() {
	conf.Init()
	util.InitLog()
	dao.MysqlInit()
	cache.RedisInit()
}
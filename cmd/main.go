package main

import (
	"fmt"
	_ "fmt"
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
	if err := util.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	conf.Init()
	util.InitLog()
	dao.MysqlInit()
	cache.RedisInit()
}

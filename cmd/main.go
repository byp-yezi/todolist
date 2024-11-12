package main

import (
	conf "todolist/config"
	"todolist/pkg/util"
)

func main() {
	loading()
}

func loading() {
	conf.Init()
	util.InitLog()
}
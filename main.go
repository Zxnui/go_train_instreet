package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	_ "go_train_instreet/routers"
	_ "go_train_instreet/task"
)

func main() {
	//定时任务
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}

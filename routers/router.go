package routers

import (
	"github.com/astaxie/beego"
	//"go_train_instreet/controllers/admin"
	"go_train_instreet/controllers/train"
)

func init() {
	beego.Router("/", &train.HomeController{})

	beego.Router("/train/select", &train.TrainController{}, "*:SelectTrain")
}

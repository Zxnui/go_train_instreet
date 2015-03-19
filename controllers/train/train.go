package train

import (
	"github.com/astaxie/beego"
)

type TrainController struct {
	beego.Controller
}

func (this *TrainController) SelectTrain() {
	this.Data["httpweb"] = Httpweb
	fromSite := this.GetString("fromSite")
	toSite := this.GetString("toSite")
	this.Data["toSite"] = toSite
	this.Data["fromSite"] = fromSite
	this.TplNames = "train.html"
}

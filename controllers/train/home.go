package train

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["httpweb"] = Httpweb
	this.TplNames = "index.html"
}

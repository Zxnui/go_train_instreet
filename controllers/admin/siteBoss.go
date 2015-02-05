package admin

import (
	"github.com/astaxie/beego"
)

type SiteBossController struct {
	beego.Controller
}

func (this *SiteBossController) GetAllSiteBoss() {

	this.Data["httpweb"] = Httpweb
	this.TplNames = "admin_site.tpl"
}

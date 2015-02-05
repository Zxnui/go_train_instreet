package admin

import (
	"github.com/astaxie/beego"
)

type SiteController struct {
	beego.Controller
}

func (this *SiteController) GetAllSite() {

	this.Data["httpweb"] = Httpweb
	this.TplNames = "admin_site.tpl"
}

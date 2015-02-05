package routers

import (
	"github.com/astaxie/beego"
	"go_train_instreet/controllers/admin"
	"go_train_instreet/controllers/train"
)

func init() {
	beego.Router("/", &train.HomeController{})
	beego.Router("/admin/site", &admin.SiteController{}, "get:GetAllSite")

	beego.Router("/admin/siteBoss", &admin.SiteBossController{}, "get:GetAllSiteBoss")
}

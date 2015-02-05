package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Cfg = beego.AppConfig

func init() {
	dbhost := beego.AppConfig.String("db_host")
	dbport := beego.AppConfig.String("db_port")
	dbuser := beego.AppConfig.String("db_user")
	dbpassword := beego.AppConfig.String("db_pass")
	dbname := beego.AppConfig.String("db_name")
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dburl)
	orm.RegisterModel(new(Train), new(Site), new(SiteBoss), new(TrainSite))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

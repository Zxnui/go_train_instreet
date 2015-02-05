package train

import (
	"fmt"
	"github.com/astaxie/beego"
)

var Httpweb string

func init() {
	httpweb := beego.AppConfig.String("httpweb")
	httpport := beego.AppConfig.String("httpport")
	fmt.Println(httpport)
	Httpweb = httpweb + ":" + httpport
}

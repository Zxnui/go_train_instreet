package task

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	//"io/ioutil"
//	//"os"
//	"strings"
//)

//func init() {
//	tk4 := toolbox.NewTask("tk4", "59 21 16 * * *", dateInfoSiteManager)
//	toolbox.AddTask("tk4", tk4)
//}
//func dateInfoSiteManager() error {
//	train := new(models.Train)
//	_, lists, err := train.QueryAll()

//	if err != nil {
//		beego.Info(err)
//		return err
//	}

//	for _, item := range lists {
//		fullName := item.Full_name
//		a := strings.Index(fullName, "(")
//		b := strings.Index(fullName, ")")
//		c := strings.Index(fullName, "-")
//		start := fullName[a+1 : c]
//		end := fullName[c+1 : b]

//		trainSiteIn(start, end, item.Id)

//	}
//	return err
//}

//func trainSiteIn(start string, end string, trainId int) {

//	var startId int
//	var endId int

//	site := new(models.Site)
//	site.Name = start
//	err := site.Read("name")
//	if err == nil {
//		startId = site.Id
//	} else {
//		_ = site.Insert()
//		_ = site.Read("name")
//		startId = site.Id
//	}

//	site = new(models.Site)
//	site.Name = end
//	err = site.Read("name")
//	if err == nil {
//		endId = site.Id
//	} else {
//		_ = site.Insert()
//		_ = site.Read("name")
//		endId = site.Id
//	}

//	trainSite := new(models.TrainSite)
//	trainSite.Site_id = startId
//	trainSite.Train_id = trainId
//	_ = trainSite.Insert()

//	trainSite = new(models.TrainSite)
//	trainSite.Site_id = endId
//	trainSite.Train_id = trainId
//	_ = trainSite.Insert()

//}

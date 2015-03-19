package task

//import (
//	//"bytes"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	//"io/ioutil"
//	//"net/http"
//)

//func init() {
//	tk2 := toolbox.NewTask("tk2", "10 06 18 * * *", dateInfoSite)
//	toolbox.AddTask("tk2", tk2)
//}

//func dateInfoSite() error {
//	beego.Info("=========获取最新的车次路径站点信息=========")

//	train := new(models.Train)
//	_, lists, err := train.QueryAll()

//	if err != nil {
//		beego.Info(err)
//		return err
//	}

//	for _, item := range lists {
//		var startName string
//		var endName string

//		trainSite := new(models.TrainSite)
//		_, trainSiteOne, _ := trainSite.QueryTwo(item.Id)

//		site := new(models.Site)
//		site.Id = trainSiteOne[0].Site_id
//		_ = site.Read("id")
//		startName = site.Data

//		site = new(models.Site)
//		site.Id = trainSiteOne[1].Site_id
//		_ = site.Read("id")
//		endName = site.Data

//		train := new(models.Train)
//		train.Id = item.Id
//		_ = train.Read("id")
//		train.Start_no = startName
//		train.End_no = endName
//		_ = train.Update("start_no", "end_no")

//	}
//	return err
//}

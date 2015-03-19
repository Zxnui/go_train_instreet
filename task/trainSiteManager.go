package task

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)

//const TRAIN_DETAIL_SITE_API_URL = "https://kyfw.12306.cn/otn/czxx/queryByTrainNo?"

//func init() {
//	tk2 := toolbox.NewTask("tk2", "10 * * * * *", dateInfoSite)
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

//		url := TRAIN_DETAIL_SITE_API_URL + "train_no=" + item.Train_no + "&from_station_telecode=" + item.Start_no + "&to_station_telecode=" + item.End_no + "&depart_date=2015-04-30"

//		response, err2 := http.Get(url)
//		if err2 != nil {
//			beego.Info(err2)
//			return err2
//		}

//		body, err1 := ioutil.ReadAll(response.Body)
//		if err1 != nil {
//			beego.Info(err1)
//			return err1
//		}
//		val := string(body)

//		a := strings.Index(val, ":[{")
//		b := strings.Index(val, "}]}")
//		y := val[a+1 : b+1]
//		li := strings.Split(y, "},{")
//		for _, siteItem := range li {
//			trainSite := new(models.TrainSite)
//			trainSite.Train_id = item.Id

//			dateli := strings.Split(siteItem, ",")

//			var siteName string
//			var startTime string
//			var stopTime string
//			var stationNo string
//			for _, dateItem := range dateli {
//				if strings.Contains(dateItem, "\"station_name") {
//					i := strings.Index(dateItem, "\":\"")

//					siteName = dateItem[i+3 : len(dateItem)-1]
//				}
//				if strings.Contains(dateItem, "start_time") {
//					i := strings.Index(dateItem, "\":\"")
//					startTime = dateItem[i+3 : len(dateItem)-1]
//				}
//				if strings.Contains(dateItem, "arrive_time") {
//					i := strings.Index(dateItem, "\":\"")
//					stopTime = dateItem[i+3 : len(dateItem)-1]

//				}
//				if strings.Contains(dateItem, "station_no") {
//					i := strings.Index(dateItem, "\":\"")
//					stationNo = dateItem[i+3 : len(dateItem)-1]
//				}
//			}
//			beego.Info(siteName)
//			beego.Info(startTime)
//			beego.Info(stopTime)
//			beego.Info(stationNo)

//			site := new(models.Site)
//			site.Name = siteName
//			err = site.Read("name")
//			if err != nil {
//				site = new(models.Site)
//				site.Name = siteName
//				_ = site.Insert()
//				trainSite.Site_id = site.Id
//			} else {
//				trainSite.Site_id = site.Id
//			}
//			trainSite.Start_time = startTime
//			trainSite.Arrive_time = stopTime
//			trainSite.Station_no = stationNo
//			_ = trainSite.Insert()

//		}
//		response.Body.Close()
//	}

//	return err
//}

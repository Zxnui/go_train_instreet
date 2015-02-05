package task

import (
	"bytes"
	"encoding/xml"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"go_train_instreet/models"
	"io/ioutil"
	"net/http"
)

const TRAIN_DETAIL_SITE_API_URL = "http://www.webxml.com.cn/WebServices/TrainTimeWebService.asmx/getDetailInfoByTrainCode?UserID=&TrainCode="

type Result struct {
	TrainDetailInfo []TrainDetailInfo
}
type TrainDetailInfo struct {
	TrainStation string
	ArriveTime   string
	StartTime    string
	KM           string
}

func init() {
	tk2 := toolbox.NewTask("tk2", "00 04 15 * * *", dateInfoSite)
	toolbox.AddTask("tk2", tk2)
}

func dateInfoSite() error {
	beego.Info("=========获取最新的车次路径站点信息=========")
	train := new(models.Train)
	allNum, lists, err := train.QueryNameAll()

	var i int64
	for i = 0; i < allNum; i++ {
		a := lists[i]["name"].(string)
		url := TRAIN_DETAIL_SITE_API_URL + a
		response, _ := http.Get(url)
		body, _ := ioutil.ReadAll(response.Body)
		num := bytes.Index(body, []byte("<getDetailInfo xmlns="))
		body = body[num:len(body)]
		var result Result
		err = xml.Unmarshal(body, &result)
		if result.TrainDetailInfo[0].ArriveTime != "http://www.webxml.com.cn" {
			for _, val := range result.TrainDetailInfo {
				train := new(models.Train)
				train.Name = a
				err = train.Read("name")
				if err == nil {
					if val.ArriveTime == "" {
						train.Start_time = val.StartTime
						_ = train.Update("start_time")
						trainSite := new(models.TrainSite)
						trainSite.Site_id = train.Start_site_id
						trainSite.Train_id = train.Id
						err = trainSite.Read("train_id", "site_id")
						if err != nil {
							trainSite = new(models.TrainSite)
							trainSite.Site_id = train.Start_site_id
							trainSite.Train_id = train.Id
							trainSite.Start_time = train.Start_time
							trainSite.Km = val.KM
							_ = trainSite.Insert()
						}
					}
					if val.StartTime == "" {
						train.End_time = val.ArriveTime
						_ = train.Update("end_time")
						trainSite := new(models.TrainSite)
						trainSite.Site_id = train.End_site_id
						trainSite.Train_id = train.Id
						err = trainSite.Read("train_id", "site_id")
						if err != nil {
							trainSite = new(models.TrainSite)
							trainSite.Site_id = train.End_site_id
							trainSite.Train_id = train.Id
							trainSite.Arrive_time = train.End_time
							trainSite.Km = val.KM
							_ = trainSite.Insert()
						}
					}
					if val.StartTime != "" && val.ArriveTime != "" {
						site := new(models.Site)
						site.Name = val.TrainStation
						err = site.Read("name")
						if err == nil {
							trainSite := new(models.TrainSite)
							trainSite.Site_id = site.Id
							trainSite.Train_id = train.Id
							err = trainSite.Read("train_id", "site_id")
							if err != nil {
								trainSite = new(models.TrainSite)
								trainSite.Site_id = site.Id
								trainSite.Train_id = train.Id
								trainSite.Start_time = val.StartTime
								trainSite.Arrive_time = val.ArriveTime
								trainSite.Km = val.KM
								_ = trainSite.Insert()
							}
						} else {
							site = new(models.Site)
							site.Name = val.TrainStation
							_ = site.Insert()

							site = new(models.Site)
							site.Name = val.TrainStation
							_ = site.Read("name")

							trainSite := new(models.TrainSite)
							trainSite.Site_id = site.Id
							trainSite.Train_id = train.Id
							err = trainSite.Read("train_id", "site_id")
							if err != nil {
								trainSite = new(models.TrainSite)
								trainSite.Site_id = site.Id
								trainSite.Train_id = train.Id
								trainSite.Start_time = val.StartTime
								trainSite.Arrive_time = val.ArriveTime
								trainSite.Km = val.KM
								_ = trainSite.Insert()
							}
						}
					}
				}
			}
		}
		response.Body.Close()
	}

	return err
}

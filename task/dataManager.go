package task

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	"io/ioutil"
//	"os"
//	"strings"
//)

//const TRAIN_FILE = "static/file/trainCar.txt"

//func init() {
//	tk1 := toolbox.NewTask("tk1", "00 15 14 * * *", dateInfo)
//	toolbox.AddTask("tk1", tk1)
//}

//func dateInfo() error {
//	beego.Info("=========timeJob begin=========")
//	fi, err := os.Open(TRAIN_FILE)
//	if err != nil {
//		beego.Info(err)
//		return err
//	}
//	defer fi.Close()

//	fd, err := ioutil.ReadAll(fi)
//	if err != nil {
//		beego.Info(err)
//		return err
//	}

//	fds := string(fd)
//	a := strings.Split(fds, ")")
//	for _, val := range a {
//		b := strings.Index(val, "(")
//		c := strings.Index(val, "-")
//		trainName := val[0:b]
//		trainStart := val[b+1 : c]
//		trainEnd := val[c+1 : len(val)]
//		insertData(trainName, trainStart, trainEnd)
//	}
//	return err
//}

//func insertData(tn string, ts string, te string) {
//	train := new(models.Train)
//	var site_start_id int
//	var site_end_id int

//	site := new(models.Site)
//	site.Name = ts
//	err := site.Read("name")
//	if err != nil {
//		err := site.Insert()
//		if err != nil {
//			beego.Info(err)
//		}
//	}

//	site = new(models.Site)
//	site.Name = te
//	err = site.Read("name")
//	if err != nil {
//		err := site.Insert()
//		if err != nil {
//			beego.Info(err)
//		}
//	}

//	site = new(models.Site)
//	site.Name = ts
//	err = site.Read("name")
//	if err == nil {
//		site_start_id = site.Id
//	}
//	site = new(models.Site)
//	site.Name = te
//	err = site.Read("name")
//	if err == nil {
//		site_end_id = site.Id
//	}

//	train.Name = tn
//	train.Start_site_id = site_start_id
//	train.End_site_id = site_end_id
//	err = train.Insert()
//	if err != nil {
//		beego.Info(err)
//	}
//}

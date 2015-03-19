package task

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	"io/ioutil"
//	"os"
//	"strings"
//)

//const TRAIN_FILE = "static/file/site.txt"

//func init() {
//	tk3 := toolbox.NewTask("tk3", "20 18 16 * * *", dateInfoSiteManager)
//	toolbox.AddTask("tk3", tk3)
//}

//func dateInfoSiteManager() error {
//	beego.Info("=========site begin=========")
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
//	a := strings.Split(fds, "@")
//	for _, val := range a {
//		if val == "" {
//			beego.Info("null")
//		} else {
//			b := strings.Index(val, "|")
//			pinyinMin := val[0:b]
//			val = val[b+1 : len(val)]

//			b = strings.Index(val, "|")
//			trainName := val[0:b]
//			val = val[b+1 : len(val)]

//			b = strings.Index(val, "|")
//			data := val[0:b]
//			val = val[b+1 : len(val)]

//			b = strings.Index(val, "|")
//			pinyinAll := val[0:b]
//			val = val[b+1 : len(val)]

//			b = strings.Index(val, "|")
//			putId := val[b+1 : len(val)]

//			insertData(trainName, data, pinyinMin, pinyinAll, putId)
//		}
//	}
//	return err
//}

//func insertData(tn string, data string, pinyinMin string, pinyinAll string, putId string) {
//	site := new(models.Site)
//	site.Name = tn
//	err := site.Read("name")
//	if err == nil {
//		site.Put_id = putId
//		site.Pinyin_min = pinyinMin
//		site.Pinyin_all = pinyinAll
//		site.Data = data
//		_ = site.Update("pinyin_min", "pinyin_all", "data", "put_id")
//	} else {
//		site.Put_id = putId
//		site.Pinyin_min = pinyinMin
//		site.Pinyin_all = pinyinAll
//		site.Data = data
//		_ = site.Insert()
//	}

//}

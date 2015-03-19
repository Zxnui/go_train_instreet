package task

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/toolbox"
//	"go_train_instreet/models"
//	"io/ioutil"
//	"os"
//	"strings"
//)

//const TRAIN_FILE = "static/file/train.txt"

//func init() {
//	tk1 := toolbox.NewTask("tk1", "30 20 16 * * *", dateInfo)
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
//	a := strings.Split(fds, "},")
//	for _, val := range a {
//		b := strings.Index(val, ":")
//		c := strings.Index(val, ")")
//		fullName := val[b+2 : c+1]
//		d := strings.Index(fullName, "(")
//		name := fullName[0:d]

//		last := val[c+1 : len(val)]
//		e := strings.Index(last, ":")
//		trainno := last[e+2 : len(last)-1]

//		err = insertData(fullName, name, trainno)
//		if err != nil {
//			beego.Info(err)
//			return err
//		}
//	}
//	return err
//}

//func insertData(tfn string, tn string, tno string) error {
//	train := new(models.Train)

//	train.Name = tn
//	train.Full_name = tfn
//	train.Train_no = tno
//	err := train.Insert()
//	return err
//}

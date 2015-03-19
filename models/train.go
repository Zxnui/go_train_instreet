package models

import (
	"github.com/astaxie/beego/orm"
)

type Train struct {
	Id        int
	Name      string
	Full_name string
	Train_no  string
	Start_no  string
	End_no    string
}

//func (m *Train) TableName() string {
//	return TableName("site")
//}

func (m *Train) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Train) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Train) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Train) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Train) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Train) QueryAll() (int64, []Train, error) {
	var lists []Train
	num, err := orm.NewOrm().Raw("select id,name,full_name,train_no,start_no,end_no from train").QueryRows(&lists)
	return num, lists, err
}

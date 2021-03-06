package models

import (
	"github.com/astaxie/beego/orm"
)

type TrainSite struct {
	Id          int
	Train_id    int
	Site_id     int
	Start_time  string
	Arrive_time string
	Km          string
	Station_no  string
}

//func (m *Site) TableName() string {
//	return TableName("site")
//}

func (m *TrainSite) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *TrainSite) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *TrainSite) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *TrainSite) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *TrainSite) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *TrainSite) QueryTwo(id int) (int64, []TrainSite, error) {
	var item []TrainSite
	num, err := orm.NewOrm().Raw("select id,train_id,site_id from train_site where train_id=?", id).QueryRows(&item)
	return num, item, err
}

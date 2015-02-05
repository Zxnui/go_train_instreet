package models

import (
	"github.com/astaxie/beego/orm"
)

type SiteBoss struct {
	Id   int
	Name string
}

//func (m *SiteBoss) TableName() string {
//	return TableName("site_boss")
//}

func (m *SiteBoss) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *SiteBoss) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *SiteBoss) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *SiteBoss) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *SiteBoss) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

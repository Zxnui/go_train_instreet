package models

import (
	"github.com/astaxie/beego/orm"
)

type Site struct {
	Id           int
	Name         string
	Site_boss_id int
	Pinyin_min   string
	Pinyin_all   string
	Data         string
	Put_id       string
}

//func (m *Site) TableName() string {
//	return TableName("site")
//}

func (m *Site) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Site) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Site) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Site) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Site) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

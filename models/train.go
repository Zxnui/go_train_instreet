package models

import (
	"github.com/astaxie/beego/orm"
)

type Train struct {
	Id            int
	Name          string
	Start_site_id int
	End_site_id   int
	Start_time    string
	End_time      string
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

func (m *Train) QueryNameAll() (int64, []orm.Params, error) {
	var lists []orm.Params
	num, err := orm.NewOrm().Raw("select name from train where id >(SELECT max(id) FROM train where start_time !='0000-00-00 00:00:00')").Values(&lists)
	return num, lists, err
}

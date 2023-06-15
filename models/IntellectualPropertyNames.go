package models

import "time"

type IntellectualPropertyNames struct {
	Id int `orm:"column(id);auto"`

	IntellectualPropertyNames string `orm:"column(intellectual_property_names);null" description:"'ip所属'"`

	Enable int `orm:"column(enable);null" description:"是否提供给前端"`

	AddTime time.Time `orm:"column(add_time);type(datetime);null" description:"创建时间"`

	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

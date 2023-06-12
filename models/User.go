package models

import "time"

type User struct {
	Id int `orm:"column(id);auto"`

	Email string `orm:"column(email);null" description:"'email'"`

	UserName string `orm:"column(user_name);null" description:"用户名"`

	PassWord string `orm:"column(password);null" description:"密码"`

	IsAdmin string `orm:"column(is_admin);null" description:"是否为管理员"`

	AddTime time.Time `orm:"column(add_time);type(datetime);null" description:"创建时间"`

	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

package models

import "time"

type KeyWord struct {
	Id int `orm:"column(id);auto"`

	KeyWord string `orm:"column(keyword);null" description:"'关键词'"`

	InfoType string `orm:"column(info_type);size(16);null" description:"片段类型(0-视频 1-图片 2-同人)"`

	Enable int `orm:"column(enable);null" description:"是否提供给前端"`

	AddTime time.Time `orm:"column(add_time);type(datetime);null" description:"创建时间"`

	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

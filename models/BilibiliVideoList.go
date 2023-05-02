package models

import "time"

type BilibiliVideoList struct {
	Id int `orm:"column(id);auto"`

	Title string `orm:"column(title);null" description:"'email'"`

	Length string `orm:"column(length);null" description:"密码"`

	Bvid string `orm:"column(bvid);null" description:"是否为管理员"`

	VideoUrl string `orm:"column(video_url);null" description:"是否为管理员"`

	Pic string `orm:"column(pic);null" description:"是否为管理员"`

	Created time.Time `orm:"column(created);type(datetime);null" description:"发布时间"`
}

package models

import "time"

type VideoList struct {
	Id int `orm:"column(id);auto"`

	VideoId string `orm:"column(video_id);null" description:"'email'"`

	VideoUrl string `orm:"column(video_url);null" description:"视频链接"`

	PublishedAt time.Time `orm:"column(published_at);type(datetime);null" description:"发布时间"`

	Title string `orm:"column(title);null" description:"视频标题"`

	Length string `orm:"column(length);null" description:"视频标题"`

	CoverOri string `orm:"column(cover_ori);null" description:"原始封面页"`
}

package models

import "time"

type BilibiliVideoList struct {
	Id int `orm:"column(id);auto"`

	Title string `orm:"column(title);null" description:"'视频标题'"`

	Length string `orm:"column(length);null" description:"视频长度"`

	VideoId string `orm:"column(video_id);null" description:"视频id"`

	VideoUrl string `orm:"column(video_url);null" description:"视频链接"`

	coverOri string `orm:"column(coverOri);null" description:"封面下载地址"`

	PublishedAt time.Time `orm:"column(published_at);type(datetime);null" description:"发布时间"`
}

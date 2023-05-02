package models

import "time"

type YoutubeVideoList struct {
	Id int `orm:"column(id);auto"`

	VideoId string `orm:"column(video_id);null" description:"'email'"`

	VideoUrl string `orm:"column(video_url);null" description:"视频链接"`

	Title string `orm:"column(title);null" description:"视频标题"`

	CoverOri string `orm:"column(cover_ori);null" description:"原始封面页"`

	CoverOss string `orm:"column(cover_oss);null" description:"另存oss封面页"`

	PublishedAt time.Time `orm:"column(published_at);type(datetime);null" description:"发布时间"`
}

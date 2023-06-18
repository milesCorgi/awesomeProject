package models

type VideoNoteInfo struct {
	Id int `orm:"column(id);auto"`

	Info string `orm:"column(info);type(text);null" description:"'详细描述'"`

	KeyWord string `orm:"column(key_word);null" description:"'关键词'"`

	TheTime string `orm:"column(the_time);null" description:"在该视频的哪个时段"`

	BiliBiliLink string `orm:"column(bili_bili_link);null" description:"b站链接"`

	BiliBiliID string `orm:"column(bili_bili_id);null" description:"b站号"`

	YoutubeLink string `orm:"column(youtube_link);null" description:"油管链接"`

	OtherLink string `orm:"column(other_link);null" description:"其余渠道的链接"`

	ImgLinks string `orm:"column(img_links);type(text);null" description:"外链图"`

	FanFictionLink string `orm:"column(fan_fiction_link);type(text);null" description:"同人相关链接"`

	InfoType string `orm:"column(info_type);size(16);null" description:"片段记录类型(0-视频 1-图片 2-同人)"`

	EditPassWord string `orm:"column(edit_password);null" description:"编辑验证码"`

	Enable int `orm:"column(enable);null" description:"是否提供给前端"`

	User string `orm:"column(user);null" description:"提交人"`

	IntellectualPropertyName string `orm:"column(intellectual_property_name);null" description:"所属ip"`

	AddTime string `orm:"column(add_time);type(datetime);null" description:"创建时间"`

	UpdateTime string `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

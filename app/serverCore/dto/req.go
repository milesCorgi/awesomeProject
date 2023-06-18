package dto

type VideoInfo struct {
	Info                     string `json:"Info" xml:"Info"`
	BiliBiliID               string `json:"BiliBiliID" xml:"BiliBiliID"`
	BiliBiliLink             string `json:"BiliBiliLink" xml:"BiliBiliLink"`
	YoutubeLink              string `json:"YoutubeLink" xml:"YoutubeLink"`
	OtherLink                string `json:"OtherLink" xml:"OtherLink"`
	FanFictionLink           string `json:"FanFictionLink" xml:"FanFictionLink"`
	InfoType                 string `json:"InfoType" xml:"InfoType"`
	KeyWord                  string `json:"KeyWord" xml:"KeyWord"`
	ImgLinks                 string `json:"ImgLinks" xml:"ImgLinks"`
	IntellectualPropertyName string `json:"IntellectualPropertyName" xml:"IntellectualPropertyName"`
	User                     string `json:"User" xml:"User"`
	EditPassWord             string `json:"EditPassWord" xml:"EditPassWord"`
}

type QueryWebVideoInfo struct {
	From     string `json:"From" xml:"From"`
	To       string `json:"To" xml:"To"`
	IfDesc   string `json:"IfDesc" xml:"IfDesc"`
	Keyword  string `json:"Keyword" xml:"Keyword"`
	SourceId int    `json:"SourceId" xml:"SourceId"`
}

type QueryKeyWord struct {
	IntellectualPropertyName string `json:"IntellectualPropertyName" xml:"IntellectualPropertyName"`
}

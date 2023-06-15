package dto

type QueryVideoInfo struct {
	Info                     string `json:"Info" xml:"Info"`
	BiliBiliID               string `json:"BiliBiliID" xml:"BiliBiliID"`
	YoutubeLink              string `json:"YoutubeLink" xml:"YoutubeLink"`
	OtherLink                string `json:"OtherLink" xml:"OtherLink"`
	InfoType                 string `json:"InfoType" xml:"InfoType"`
	KeyWord                  string `json:"KeyWord" xml:"KeyWord"`
	IntellectualPropertyName string `json:"IntellectualPropertyName" xml:"IntellectualPropertyName"`
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

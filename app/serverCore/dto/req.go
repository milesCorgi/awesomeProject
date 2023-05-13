package dto

type QueryAInfo struct {
	Info        string `json:"Info" xml:"Info"`
	BiliBiliID  string `json:"BiliBiliID" xml:"BiliBiliID"`
	YoutubeLink string `json:"YoutubeLink" xml:"YoutubeLink"`
	OtherLink   string `json:"OtherLink" xml:"OtherLink"`
	InfoType    string `json:"InfoType" xml:"InfoType"`
	KeyWord     string `json:"KeyWord" xml:"KeyWord"`
}

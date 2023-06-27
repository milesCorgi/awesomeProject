package dto

type CacheInfo struct {
	Key           string `json:"Key" xml:"Key"`
	Value         string `json:"Value" xml:"Value"`
	ExpiredSecond string `json:"ExpiredSecond" xml:"ExpiredSecond"`
}

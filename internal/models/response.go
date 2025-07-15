package models

type Response struct {
	LongUrl  string `json:"src"`
	ShortUrl string `json:"res"`
}

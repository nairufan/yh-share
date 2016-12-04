package util

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type urlResponse struct {
	ErrCode  int       `json:"errcode"`
	ErrMsg   string        `json:"errmsg"`
	ShortUrl string        `json:"short_url"`
}

func GetUrl(url string, token string) string {
	shortUrl := beego.AppConfig.String("wechat.shortUrl")
	shortUrl = fmt.Sprintf(shortUrl, token)
	resp, err := http.Get(shortUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	beego.Info(string(body))
	response := &urlResponse{}
	json.Unmarshal(body, response)
	if response.ErrCode != 0 {
		return ""
	}
	return response.ShortUrl
}

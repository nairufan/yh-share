package util

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const Delta = 1200

type Token struct {
	AccessToken string       `json:"access_token"`
	ExpiresIn   int64        `json:"expires_in"`
}

func GetToken() *Token {
	appId := beego.AppConfig.String("wechat.appId")
	secret := beego.AppConfig.String("wechat.appSecret")
	authUrl := beego.AppConfig.String("wechat.accessTokenUrl")
	authUrl = fmt.Sprintf(authUrl, appId, secret)
	resp, err := http.Get(authUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	beego.Info(string(body))
	response := &Token{}
	json.Unmarshal(body, response)
	return response
}

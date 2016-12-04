package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/bluesuncorp/validator.v5"
	"github.com/nairufan/yh-share/util"
	"time"
)

const (
	UserID = "userId"
	Excel = "excel"
	Token = "token"
	Expire = "expire"
)

var validate = validator.New("validate", validator.BakedInValidators)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) SetUserId(id string) {
	b.SetSession(UserID, id)
}

func (b *BaseController) GetUserId() string {
	uid := b.GetSession(UserID)
	if uid == nil {
		return ""
	}
	return uid.(string)
}

func (b *BaseController) SetExcel(data [][]string) {
	b.SetSession(Excel, data)
}

func (b *BaseController) getExcel() [][]string {
	excel := b.GetSession(Excel)
	if excel == nil {
		return nil
	}
	return excel.([][]string)
}

func (b *BaseController) ClearExcel() [][]string {
	excel := b.getExcel()
	b.DelSession(Excel)
	return excel
}

func (b *BaseController) SetToken(token string) {
	b.SetSession(Token, token)
}

func (b *BaseController) GetToken() string {
	token := b.GetSession(Token)
	expire := b.GetExpire()
	now := time.Now().Unix()
	if token == nil || now > expire {
		newToken := util.GetToken()
		b.SetToken(newToken.AccessToken)
		b.SetExpire(newToken.ExpiresIn)
		token = newToken.AccessToken
	}
	return token.(string)
}

func (b *BaseController) SetExpire(expire int64) {
	b.SetSession(Expire, expire + time.Now().Unix() - util.Delta)
}

func (b *BaseController) GetExpire() int64 {
	expire := b.GetSession(Expire)
	if expire == nil {
		return 0
	}
	return expire.(int64)
}
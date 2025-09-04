package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/internal/model"
)

type WxLoginReq struct {
	g.Meta `path:"/third/login" method:"post" summary:"微信登录" tags:"Third Party"`
	Code   string `json:"code" v:"required#微信登录code不能为空" dc:"微信登录code"`
}

type WxLoginRes model.WechatLoginOutput

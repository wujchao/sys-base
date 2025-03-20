package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaIndexReq struct {
	g.Meta `path:"/captcha" method:"get" rules:"SysCaptcha" auth:"open" tags:"Log in" summary:"Get the default verification code"`
}

type CaptchaIndexRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key" dc:"key"`
	Img    string `json:"img" dc:"picture"`
}

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/internal/model"
)

type LoginDoReq struct {
	g.Meta    `path:"/login" method:"post" rules:"SysLogin" auth:"open" summary:"Execute login request" tags:"Log in"`
	UserName  string `json:"username" v:"required#Please enter your account number"   dc:"username"`
	Password  string `json:"password" v:"required#Please enter your password"   dc:"password"`
	Captcha   string `json:"captcha" dc:"Verification Code"`
	VerifyKey string `json:"verifyKey"`
	Module    string `json:"module"`
}
type LoginDoRes struct {
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	IsChangePwd int                 `json:"isChangePwd" dc:"Do you need to change the password? 1 Yes 2 No"`
}

type LoginOutReq struct {
	g.Meta `path:"/loginOut" method:"post" rules:"SysLogout" auth:"open" summary:"Sign out" tags:"Log in"`
}
type LoginOutRes struct {
}

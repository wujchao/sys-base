package model

// WechatLoginResponse 定义微信接口返回结果的结构体
type WechatLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type WechatLoginOutput struct {
	*ThirdPartyLoginOutput
}

type ThirdPartyLoginInput struct {
	ThirdPartyName   string `json:"thirdPartyName"   orm:"third_party_name"    description:"第三方平台名称"`      // 第三方平台名称
	ThirdPartyUserId string `json:"thirdPartyUserId" orm:"third_party_user_id" description:"第三方平台的用户唯一标识"` // 第三方平台的用户唯一标识
	AccessToken      string `json:"accessToken"      orm:"access_token"        description:"第三方平台的访问令牌"`   // 第三方平台的访问令牌
}

type ThirdPartyLoginOutput struct {
	UserInfo    *LoginUserRes `json:"userInfo"`
	Token       string        `json:"token" dc:"token"`
	IsChangePwd int           `json:"isChangePwd" dc:"Do you need to change the password? 1 Yes 0|2 No"`
}

type ThirdPartyRegisterInput struct {
	UserName         string `json:"userName"       orm:"user_name"       description:"用户名"` // 用户名
	ThirdPartyUserId string `json:"ThirdPartyUserId" dc:"第三方用户ID"`
	AccessToken      string `json:"accessToken"      orm:"access_token"        description:"第三方平台的访问令牌"` // 第三方平台的访问令牌
	UserTypes        int    `json:"userTypes"      orm:"user_types"      description:"1 管理，2 企业，9系统"`    // 1 管理，2 企业，3 用户，9系统
}

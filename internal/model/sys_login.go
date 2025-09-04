package model

type LoginOutput struct {
	UserInfo    *LoginUserRes `json:"userInfo"`
	Token       string        `json:"token"`
	IsChangePwd int           `json:"isChangePwd" dc:"Do you need to change the password? 1 Yes 0|2 No"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysThirdPartyUser is the golang structure for table sys_third_party_user.
type SysThirdPartyUser struct {
	Id               int64       `json:"id"               orm:"id"                  description:"关联记录唯一标识"`             // 关联记录唯一标识
	UserId           string      `json:"userId"           orm:"user_id"             description:"用户 ID，关联 users 表的 id"` // 用户 ID，关联 users 表的 id
	ThirdPartyName   string      `json:"thirdPartyName"   orm:"third_party_name"    description:"第三方平台名称"`              // 第三方平台名称
	ThirdPartyUserId string      `json:"thirdPartyUserId" orm:"third_party_user_id" description:"第三方平台的用户唯一标识"`         // 第三方平台的用户唯一标识
	AccessToken      string      `json:"accessToken"      orm:"access_token"        description:"第三方平台的访问令牌"`           // 第三方平台的访问令牌
	RefreshToken     string      `json:"refreshToken"     orm:"refresh_token"       description:"第三方平台的刷新令牌"`           // 第三方平台的刷新令牌
	ExpiresAt        *gtime.Time `json:"expiresAt"        orm:"expires_at"          description:"访问令牌过期时间"`             // 访问令牌过期时间
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"          description:"关联记录创建时间"`             // 关联记录创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"          description:"关联记录更新时间"`             // 关联记录更新时间
}

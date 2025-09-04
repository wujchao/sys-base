// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysThirdPartyUser is the golang structure of table sys_third_party_user for DAO operations like Where/Data.
type SysThirdPartyUser struct {
	g.Meta           `orm:"table:sys_third_party_user, do:true"`
	Id               interface{} // 关联记录唯一标识
	UserId           interface{} // 用户 ID，关联 users 表的 id
	ThirdPartyName   interface{} // 第三方平台名称
	ThirdPartyUserId interface{} // 第三方平台的用户唯一标识
	AccessToken      interface{} // 第三方平台的访问令牌
	RefreshToken     interface{} // 第三方平台的刷新令牌
	ExpiresAt        *gtime.Time // 访问令牌过期时间
	CreatedAt        *gtime.Time // 关联记录创建时间
	UpdatedAt        *gtime.Time // 关联记录更新时间
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUsers is the golang structure of table sys_users for DAO operations like Where/Data.
type SysUsers struct {
	g.Meta         `orm:"table:sys_users, do:true"`
	Id             interface{} // ID
	OrgId          interface{} // 企业ID
	UserName       interface{} // 用户名
	UserTypes      interface{} // 1 管理，2 企业，9系统
	Mobile         interface{} // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname   interface{} // 用户昵称
	UserPassword   interface{} // 登录密码
	UserSalt       interface{} // 加密盐
	Sex            interface{} // 性别0保密1男2女
	UserEmail      interface{} // 用户登录邮箱
	Avatar         interface{} // 用户头像
	DeptId         interface{} // 部门id
	Remark         interface{} // 备注
	IsAdmin        interface{} // 是否后台管理员 1 是  2 否
	Address        interface{} // 联系地址
	LastLoginIp    interface{} // 最后登录ip
	LastLoginTime  *gtime.Time // 最后登录时间
	Status         interface{} // 用户状态; 1:正常,2:禁用,3:未验证
	EffectiveDate  *gtime.Time // 生效日期
	ExpirationDate *gtime.Time // 过期日期
	CreatedBy      interface{} // 创建者
	CreatedAt      *gtime.Time // 创建日期
	UpdatedAt      *gtime.Time // 修改日期
	DeletedAt      interface{} // 删除时间
}

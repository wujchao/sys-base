// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRules is the golang structure of table sys_rules for DAO operations like Where/Data.
type SysRules struct {
	g.Meta     `orm:"table:sys_rules, do:true"`
	Id         interface{} // 权限标识 全局唯一
	ParentId   interface{} // 父ID
	Name       interface{} // 名称
	Apis       interface{} // api地址列表
	Menus      interface{} // 权限对应显示的菜单
	FrontIds   interface{} // 前置权限
	ListOrder  interface{} // 排序，从小到大
	ModuleType interface{} // 类型 1全部 2 管理 3 企业
	AuthType   interface{} // 授权类型，1 授权，2 登录，3 公开
	Status     interface{} // 状态
	Desc       interface{} // 说明
	Remark     interface{} // 备注
}

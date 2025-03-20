// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleRules is the golang structure of table sys_role_rules for DAO operations like Where/Data.
type SysRoleRules struct {
	g.Meta  `orm:"table:sys_role_rules, do:true"`
	RoleId  interface{} // 角色ID
	RulesId interface{} // 规则标识
}

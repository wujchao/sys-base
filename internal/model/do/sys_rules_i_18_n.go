// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRulesI18N is the golang structure of table sys_rules_i18n for DAO operations like Where/Data.
type SysRulesI18N struct {
	g.Meta   `orm:"table:sys_rules_i18n, do:true"`
	Id       interface{} //
	LangCode interface{} // 语言代码
	Name     interface{} // 名称
	Desc     interface{} // 说明
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSetting is the golang structure of table sys_setting for DAO operations like Where/Data.
type SysSetting struct {
	g.Meta       `orm:"table:sys_setting, do:true"`
	Id           interface{} // 系统设置key
	SettingGroup interface{} // 设置分组
	Value        interface{} // 设置值
	Name         interface{} // 设置名称
	Description  interface{} // 描述
	Type         interface{} // 1系统内置2否
	Status       interface{} // 状态1启用2禁用
	SortList     interface{} // 排序
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}

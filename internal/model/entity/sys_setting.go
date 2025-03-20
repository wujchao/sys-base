// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSetting is the golang structure for table sys_setting.
type SysSetting struct {
	Id           string      `json:"id"           orm:"id"            description:"系统设置key"`  // 系统设置key
	SettingGroup string      `json:"settingGroup" orm:"setting_group" description:"设置分组"`     // 设置分组
	Value        string      `json:"value"        orm:"value"         description:"设置值"`      // 设置值
	Name         string      `json:"name"         orm:"name"          description:"设置名称"`     // 设置名称
	Description  string      `json:"description"  orm:"description"   description:"描述"`       // 描述
	Type         int         `json:"type"         orm:"type"          description:"1系统内置2否"`  // 1系统内置2否
	Status       int         `json:"status"       orm:"status"        description:"状态1启用2禁用"` // 状态1启用2禁用
	SortList     int         `json:"sortList"     orm:"sort_list"     description:"排序"`       // 排序
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`         //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`         //
}

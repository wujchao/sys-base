// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id         string      `json:"id"         orm:"id"           description:"菜单标识，唯一"`             // 菜单标识，唯一
	ParentId   string      `json:"parentId"   orm:"parent_id"    description:"父ID"`                 // 父ID
	Title      string      `json:"title"      orm:"title"        description:"菜单名称"`                // 菜单名称
	Icon       string      `json:"icon"       orm:"icon"         description:"图标"`                  // 图标
	Affix      uint        `json:"affix"      orm:"affix"        description:"1 固定在tag"`            // 1 固定在tag
	NoTagsView uint        `json:"noTagsView" orm:"no_tags_view" description:"1 不显示在tag"`           // 1 不显示在tag
	Condition  string      `json:"condition"  orm:"condition"    description:"条件"`                  // 条件
	Remark     string      `json:"remark"     orm:"remark"       description:"备注"`                  // 备注
	MenuType   uint        `json:"menuType"   orm:"menu_type"    description:"类型 1目录 2菜单 3按钮"`      // 类型 1目录 2菜单 3按钮
	Weigh      int         `json:"weigh"      orm:"weigh"        description:"权重"`                  // 权重
	Hidden     uint        `json:"hidden"     orm:"hidden"       description:"1 不显示在菜单栏"`           // 1 不显示在菜单栏
	IsHide     uint        `json:"isHide"     orm:"is_hide"      description:"是否隐藏 1是 2否"`          // 是否隐藏 1是 2否
	Path       string      `json:"path"       orm:"path"         description:"路由地址"`                // 路由地址
	Component  string      `json:"component"  orm:"component"    description:"指向组件View"`            // 指向组件View
	ModuleType uint        `json:"moduleType" orm:"module_type"  description:"1 全部 2 管理 3 企业 0 系统"` // 1 全部 2 管理 3 企业 0 系统
	IsCached   uint        `json:"isCached"   orm:"is_cached"    description:"是否缓存"`                // 是否缓存
	Redirect   string      `json:"redirect"   orm:"redirect"     description:"路由重定向地址"`             // 路由重定向地址
	Open       uint        `json:"open"       orm:"open"         description:"1公开菜单"`               // 1公开菜单
	Status     int         `json:"status"     orm:"status"       description:"状态 1启用 2停用"`          // 状态 1启用 2停用
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`                // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"更新时间"`                // 更新时间
	DeletedAt  uint        `json:"deletedAt"  orm:"deleted_at"   description:"删除时间"`                // 删除时间
}

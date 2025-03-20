// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	Id         interface{} // 菜单标识，唯一
	ParentId   interface{} // 父ID
	Title      interface{} // 菜单名称
	Icon       interface{} // 图标
	Affix      interface{} // 1 固定在tag
	NoTagsView interface{} // 1 不显示在tag
	Condition  interface{} // 条件
	Remark     interface{} // 备注
	MenuType   interface{} // 类型 1目录 2菜单 3按钮
	Weigh      interface{} // 权重
	Hidden     interface{} // 1 不显示在菜单栏
	IsHide     interface{} // 是否隐藏 1是 2否
	Path       interface{} // 路由地址
	Component  interface{} // 指向组件View
	ModuleType interface{} // 1 全部 2 管理 3 企业 0 系统
	IsCached   interface{} // 是否缓存
	Redirect   interface{} // 路由重定向地址
	Open       interface{} // 1公开菜单
	Status     interface{} // 状态 1启用 2停用
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  interface{} // 删除时间
}

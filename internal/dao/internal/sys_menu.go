// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for table sys_menu.
type SysMenuColumns struct {
	Id         string // 菜单标识，唯一
	ParentId   string // 父ID
	Title      string // 菜单名称
	Icon       string // 图标
	Affix      string // 1 固定在tag
	NoTagsView string // 1 不显示在tag
	Condition  string // 条件
	Remark     string // 备注
	MenuType   string // 类型 1目录 2菜单 3按钮
	Weigh      string // 权重
	Hidden     string // 1 不显示在菜单栏
	IsHide     string // 是否隐藏 1是 2否
	Path       string // 路由地址
	Component  string // 指向组件View
	ModuleType string // 1 全部 2 管理 3 企业 0 系统
	IsCached   string // 是否缓存
	Redirect   string // 路由重定向地址
	Open       string // 1公开菜单
	Status     string // 状态 1启用 2停用
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// sysMenuColumns holds the columns for table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Title:      "title",
	Icon:       "icon",
	Affix:      "affix",
	NoTagsView: "no_tags_view",
	Condition:  "condition",
	Remark:     "remark",
	MenuType:   "menu_type",
	Weigh:      "weigh",
	Hidden:     "hidden",
	IsHide:     "is_hide",
	Path:       "path",
	Component:  "component",
	ModuleType: "module_type",
	IsCached:   "is_cached",
	Redirect:   "redirect",
	Open:       "open",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

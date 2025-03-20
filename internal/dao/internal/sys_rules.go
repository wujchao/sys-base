// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRulesDao is the data access object for table sys_rules.
type SysRulesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysRulesColumns // columns contains all the column names of Table for convenient usage.
}

// SysRulesColumns defines and stores column names for table sys_rules.
type SysRulesColumns struct {
	Id         string // 权限标识 全局唯一
	ParentId   string // 父ID
	Name       string // 名称
	Apis       string // api地址列表
	Menus      string // 权限对应显示的菜单
	FrontIds   string // 前置权限
	ListOrder  string // 排序，从小到大
	ModuleType string // 类型 1全部 2 管理 3 企业
	AuthType   string // 授权类型，1 授权，2 登录，3 公开
	Status     string // 状态
	Desc       string // 说明
	Remark     string // 备注
}

// sysRulesColumns holds the columns for table sys_rules.
var sysRulesColumns = SysRulesColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Name:       "name",
	Apis:       "apis",
	Menus:      "menus",
	FrontIds:   "front_ids",
	ListOrder:  "list_order",
	ModuleType: "module_type",
	AuthType:   "auth_type",
	Status:     "status",
	Desc:       "desc",
	Remark:     "remark",
}

// NewSysRulesDao creates and returns a new DAO object for table data access.
func NewSysRulesDao() *SysRulesDao {
	return &SysRulesDao{
		group:   "default",
		table:   "sys_rules",
		columns: sysRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRulesDao) Columns() SysRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleRulesDao is the data access object for table sys_role_rules.
type SysRoleRulesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysRoleRulesColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleRulesColumns defines and stores column names for table sys_role_rules.
type SysRoleRulesColumns struct {
	RoleId  string // 角色ID
	RulesId string // 规则标识
}

// sysRoleRulesColumns holds the columns for table sys_role_rules.
var sysRoleRulesColumns = SysRoleRulesColumns{
	RoleId:  "role_id",
	RulesId: "rules_id",
}

// NewSysRoleRulesDao creates and returns a new DAO object for table data access.
func NewSysRoleRulesDao() *SysRoleRulesDao {
	return &SysRoleRulesDao{
		group:   "default",
		table:   "sys_role_rules",
		columns: sysRoleRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleRulesDao) Columns() SysRoleRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRulesI18NDao is the data access object for table sys_rules_i18n.
type SysRulesI18NDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysRulesI18NColumns // columns contains all the column names of Table for convenient usage.
}

// SysRulesI18NColumns defines and stores column names for table sys_rules_i18n.
type SysRulesI18NColumns struct {
	Id       string //
	LangCode string // 语言代码
	Name     string // 名称
	Desc     string // 说明
}

// sysRulesI18NColumns holds the columns for table sys_rules_i18n.
var sysRulesI18NColumns = SysRulesI18NColumns{
	Id:       "id",
	LangCode: "lang_code",
	Name:     "name",
	Desc:     "desc",
}

// NewSysRulesI18NDao creates and returns a new DAO object for table data access.
func NewSysRulesI18NDao() *SysRulesI18NDao {
	return &SysRulesI18NDao{
		group:   "default",
		table:   "sys_rules_i18n",
		columns: sysRulesI18NColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRulesI18NDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRulesI18NDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRulesI18NDao) Columns() SysRulesI18NColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRulesI18NDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRulesI18NDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRulesI18NDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDocDao is the data access object for table sys_doc.
type SysDocDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysDocColumns // columns contains all the column names of Table for convenient usage.
}

// SysDocColumns defines and stores column names for table sys_doc.
type SysDocColumns struct {
	Id          string //
	Name        string // 名称
	Key         string // 标识
	Title       string // 标题
	Desc        string // 简介
	Content     string // 内容，当内容为文件时，存储文件路径及文件类型的json
	ContentType string // 1 html 2 markdowm 3 json
	Status      string // 状态 1 正常 2 禁用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// sysDocColumns holds the columns for table sys_doc.
var sysDocColumns = SysDocColumns{
	Id:          "id",
	Name:        "name",
	Key:         "key",
	Title:       "title",
	Desc:        "desc",
	Content:     "content",
	ContentType: "content_type",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysDocDao creates and returns a new DAO object for table data access.
func NewSysDocDao() *SysDocDao {
	return &SysDocDao{
		group:   "default",
		table:   "sys_doc",
		columns: sysDocColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDocDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDocDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDocDao) Columns() SysDocColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDocDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDocDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDocDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

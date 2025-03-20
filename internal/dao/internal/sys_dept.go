// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeptDao is the data access object for table sys_dept.
type SysDeptDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysDeptColumns // columns contains all the column names of Table for convenient usage.
}

// SysDeptColumns defines and stores column names for table sys_dept.
type SysDeptColumns struct {
	Id        string // 部门id
	OrgId     string // 企业ID
	ParentId  string // 父部门id
	Path      string // 路径
	DeptName  string // 部门名称
	ListOrder string // 排序
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	Status    string // 部门状态（0停用 1正常）
	CreatedAt string // 创建时间
	CreatedBy string // 创建人
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// sysDeptColumns holds the columns for table sys_dept.
var sysDeptColumns = SysDeptColumns{
	Id:        "id",
	OrgId:     "org_id",
	ParentId:  "parent_id",
	Path:      "path",
	DeptName:  "dept_name",
	ListOrder: "list_order",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSysDeptDao creates and returns a new DAO object for table data access.
func NewSysDeptDao() *SysDeptDao {
	return &SysDeptDao{
		group:   "default",
		table:   "sys_dept",
		columns: sysDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDeptDao) Columns() SysDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

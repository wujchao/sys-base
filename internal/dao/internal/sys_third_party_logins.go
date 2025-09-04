// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysThirdPartyLoginsDao is the data access object for table sys_third_party_logins.
type SysThirdPartyLoginsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns SysThirdPartyLoginsColumns // columns contains all the column names of Table for convenient usage.
}

// SysThirdPartyLoginsColumns defines and stores column names for table sys_third_party_logins.
type SysThirdPartyLoginsColumns struct {
	Id               string // 关联记录唯一标识
	UserId           string // 用户 ID，关联 users 表的 id
	ThirdPartyName   string // 第三方平台名称
	ThirdPartyUserId string // 第三方平台的用户唯一标识
	AccessToken      string // 第三方平台的访问令牌
	RefreshToken     string // 第三方平台的刷新令牌
	ExpiresAt        string // 访问令牌过期时间
	CreatedAt        string // 关联记录创建时间
	UpdatedAt        string // 关联记录更新时间
}

// sysThirdPartyLoginsColumns holds the columns for table sys_third_party_logins.
var sysThirdPartyLoginsColumns = SysThirdPartyLoginsColumns{
	Id:               "id",
	UserId:           "user_id",
	ThirdPartyName:   "third_party_name",
	ThirdPartyUserId: "third_party_user_id",
	AccessToken:      "access_token",
	RefreshToken:     "refresh_token",
	ExpiresAt:        "expires_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewSysThirdPartyLoginsDao creates and returns a new DAO object for table data access.
func NewSysThirdPartyLoginsDao() *SysThirdPartyLoginsDao {
	return &SysThirdPartyLoginsDao{
		group:   "default",
		table:   "sys_third_party_logins",
		columns: sysThirdPartyLoginsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysThirdPartyLoginsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysThirdPartyLoginsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysThirdPartyLoginsDao) Columns() SysThirdPartyLoginsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysThirdPartyLoginsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysThirdPartyLoginsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysThirdPartyLoginsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

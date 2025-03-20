// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUsersDao is the data access object for table sys_users.
type SysUsersDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysUsersColumns // columns contains all the column names of Table for convenient usage.
}

// SysUsersColumns defines and stores column names for table sys_users.
type SysUsersColumns struct {
	Id             string // ID
	OrgId          string // 企业ID
	UserName       string // 用户名
	UserTypes      string // 1 管理，2 企业，9系统
	Mobile         string // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname   string // 用户昵称
	UserPassword   string // 登录密码
	UserSalt       string // 加密盐
	Sex            string // 性别0保密1男2女
	UserEmail      string // 用户登录邮箱
	Avatar         string // 用户头像
	DeptId         string // 部门id
	Remark         string // 备注
	IsAdmin        string // 是否后台管理员 1 是  2 否
	Address        string // 联系地址
	LastLoginIp    string // 最后登录ip
	LastLoginTime  string // 最后登录时间
	Status         string // 用户状态; 1:正常,2:禁用,3:未验证
	EffectiveDate  string // 生效日期
	ExpirationDate string // 过期日期
	CreatedBy      string // 创建者
	CreatedAt      string // 创建日期
	UpdatedAt      string // 修改日期
	DeletedAt      string // 删除时间
}

// sysUsersColumns holds the columns for table sys_users.
var sysUsersColumns = SysUsersColumns{
	Id:             "id",
	OrgId:          "org_id",
	UserName:       "user_name",
	UserTypes:      "user_types",
	Mobile:         "mobile",
	UserNickname:   "user_nickname",
	UserPassword:   "user_password",
	UserSalt:       "user_salt",
	Sex:            "sex",
	UserEmail:      "user_email",
	Avatar:         "avatar",
	DeptId:         "dept_id",
	Remark:         "remark",
	IsAdmin:        "is_admin",
	Address:        "address",
	LastLoginIp:    "last_login_ip",
	LastLoginTime:  "last_login_time",
	Status:         "status",
	EffectiveDate:  "effective_date",
	ExpirationDate: "expiration_date",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewSysUsersDao creates and returns a new DAO object for table data access.
func NewSysUsersDao() *SysUsersDao {
	return &SysUsersDao{
		group:   "default",
		table:   "sys_users",
		columns: sysUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUsersDao) Columns() SysUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

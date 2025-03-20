// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta    `orm:"table:sys_role, do:true"`
	Id        interface{} // ID
	OrgId     interface{} // 企业ID
	DeptId    interface{} // 部门ID
	ParentId  interface{} // 父ID
	RoleType  interface{} // 类型，1管理，2企业，9系统
	Path      interface{} // 路径
	ListOrder interface{} // 排序
	Name      interface{} // 角色名称
	DataScope interface{} // 数据范围（1：全部数据权限 2：本部门及以下数据权限  3：本部门数据权限 4：自定数据权限 ）
	Remark    interface{} // 备注
	Status    interface{} // 状态;0:禁用;1:正常
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建日期
	UpdatedAt *gtime.Time // 修改日期
	DeletedAt interface{} // 删除时间
}

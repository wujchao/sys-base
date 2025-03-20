// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta    `orm:"table:sys_dept, do:true"`
	Id        interface{} // 部门id
	OrgId     interface{} // 企业ID
	ParentId  interface{} // 父部门id
	Path      interface{} // 路径
	DeptName  interface{} // 部门名称
	ListOrder interface{} // 排序
	Leader    interface{} // 负责人
	Phone     interface{} // 联系电话
	Email     interface{} // 邮箱
	Status    interface{} // 部门状态（0停用 1正常）
	CreatedAt *gtime.Time // 创建时间
	CreatedBy interface{} // 创建人
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt interface{} // 删除时间
}

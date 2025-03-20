// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	Id        string      `json:"id"        orm:"id"         description:"部门id"`          // 部门id
	OrgId     string      `json:"orgId"     orm:"org_id"     description:"企业ID"`          // 企业ID
	ParentId  string      `json:"parentId"  orm:"parent_id"  description:"父部门id"`         // 父部门id
	Path      string      `json:"path"      orm:"path"       description:"路径"`            // 路径
	DeptName  string      `json:"deptName"  orm:"dept_name"  description:"部门名称"`          // 部门名称
	ListOrder int         `json:"listOrder" orm:"list_order" description:"排序"`            // 排序
	Leader    string      `json:"leader"    orm:"leader"     description:"负责人"`           // 负责人
	Phone     string      `json:"phone"     orm:"phone"      description:"联系电话"`          // 联系电话
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`            // 邮箱
	Status    uint        `json:"status"    orm:"status"     description:"部门状态（0停用 1正常）"` // 部门状态（0停用 1正常）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`          // 创建时间
	CreatedBy string      `json:"createdBy" orm:"created_by" description:"创建人"`           // 创建人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`          // 修改时间
	DeletedAt uint        `json:"deletedAt" orm:"deleted_at" description:"删除时间"`          // 删除时间
}

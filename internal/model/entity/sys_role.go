// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        string      `json:"id"        orm:"id"         description:"ID"`                                               // ID
	OrgId     string      `json:"orgId"     orm:"org_id"     description:"企业ID"`                                             // 企业ID
	DeptId    string      `json:"deptId"    orm:"dept_id"    description:"部门ID"`                                             // 部门ID
	ParentId  string      `json:"parentId"  orm:"parent_id"  description:"父ID"`                                              // 父ID
	RoleType  uint        `json:"roleType"  orm:"role_type"  description:"类型，1管理，2企业，9系统"`                                   // 类型，1管理，2企业，9系统
	Path      string      `json:"path"      orm:"path"       description:"路径"`                                               // 路径
	ListOrder uint        `json:"listOrder" orm:"list_order" description:"排序"`                                               // 排序
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`                                             // 角色名称
	DataScope uint        `json:"dataScope" orm:"data_scope" description:"数据范围（1：全部数据权限 2：本部门及以下数据权限  3：本部门数据权限 4：自定数据权限 ）"` // 数据范围（1：全部数据权限 2：本部门及以下数据权限  3：本部门数据权限 4：自定数据权限 ）
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                                               // 备注
	Status    uint        `json:"status"    orm:"status"     description:"状态;0:禁用;1:正常"`                                     // 状态;0:禁用;1:正常
	CreatedBy uint        `json:"createdBy" orm:"created_by" description:"创建者"`                                              // 创建者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建日期"`                                             // 创建日期
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改日期"`                                             // 修改日期
	DeletedAt uint        `json:"deletedAt" orm:"deleted_at" description:"删除时间"`                                             // 删除时间
}

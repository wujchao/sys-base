package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RoleListReqParam struct {
	OrgId  string `json:"orgId" dc:"企业ID"`
	Name   string `json:"name" dc:"角色名称"`
	Status uint   `json:"status" dc:"状态"`
	*PaginationInput
}

type SysRoleOutput struct {
	g.Meta    `orm:"table:sys_role, do:true"`
	Id        string      `json:"id"        orm:"id"         ` // ID
	OrgId     string      `json:"orgId"     orm:"org_id"     ` // 企业ID
	DeptId    string      `json:"deptId"    orm:"dept_id"    ` // 部门ID
	ParentId  string      `json:"parentId"  orm:"parent_id"  ` // 父ID
	RoleType  uint        `json:"roleType"  orm:"role_type"  ` // 类型，1管理，2企业，9系统
	Path      string      `json:"path"      orm:"path"       ` // 路径
	ListOrder uint        `json:"listOrder" orm:"list_order" ` // 排序
	Name      string      `json:"name"      orm:"name"       ` // 角色名称
	DataScope uint        `json:"dataScope" orm:"data_scope" ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	Status    uint        `json:"status"    orm:"status"     ` // 状态;0:禁用;1:正常
	CreatedBy string      `json:"createdBy" orm:"created_by" ` // 创建者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建日期
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改日期
}

type SysRoleOutputWithUser struct {
	*SysRoleOutput
	UserId string `json:"user_id"`
}

type SysRoleListOutput struct {
	List []*SysRoleOutput
	PaginationOutput
}

type SysRoleInput struct {
	OrgId     string   `json:"orgId"     orm:"org_id"     ` // 企业ID
	DeptId    string   `json:"deptId"    orm:"dept_id"    ` // 部门ID
	ParentId  string   `json:"parentId"  orm:"parent_id"  ` // 父ID
	ListOrder uint     `json:"listOrder" orm:"list_order" ` // 排序
	Name      string   `json:"name"      orm:"name"       ` // 角色名称
	DataScope uint     `json:"dataScope" orm:"data_scope" ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Remark    string   `json:"remark"    orm:"remark"     ` // 备注
	Status    uint     `json:"status"    orm:"status"     ` // 状态;0:禁用;1:正常
	Rules     []string `json:"rules" dc:"权限"`
}

type SysRoleEditInput struct {
	Id        string   `json:"id"`
	ParentId  string   `json:"parentId"  orm:"parent_id"  ` // 父ID
	ListOrder uint     `json:"listOrder" orm:"list_order" ` // 排序
	Name      string   `json:"name"      orm:"name"       ` // 角色名称
	DataScope uint     `json:"dataScope" orm:"data_scope" ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Remark    string   `json:"remark"    orm:"remark"     ` // 备注
	Status    uint     `json:"status"    orm:"status"     ` // 状态;0:禁用;1:正常
	Rules     []string `json:"rules" dc:"权限"`
}

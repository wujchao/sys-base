// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure for table sys_organization.
type SysOrganization struct {
	Id             string      `json:"id"             orm:"id"              description:"ID"`          // ID
	ParentId       string      `json:"parentId"       orm:"parent_id"       description:"父id"`         // 父id
	Path           string      `json:"path"           orm:"path"            description:"级别路径"`        // 级别路径
	Name           string      `json:"name"           orm:"name"            description:"名称"`          // 名称
	Number         string      `json:"number"         orm:"number"          description:"编号"`          // 编号
	ListOrder      int         `json:"listOrder"      orm:"list_order"      description:"排序"`          // 排序
	Leader         string      `json:"leader"         orm:"leader"          description:"负责人"`         // 负责人
	Phone          string      `json:"phone"          orm:"phone"           description:"联系电话"`        // 联系电话
	Email          string      `json:"email"          orm:"email"           description:"邮箱"`          // 邮箱
	Status         uint        `json:"status"         orm:"status"          description:"状态（0停用 1正常）"` // 状态（0停用 1正常）
	EffectiveDate  *gtime.Time `json:"effectiveDate"  orm:"effective_date"  description:"生效日期"`        // 生效日期
	ExpirationDate *gtime.Time `json:"expirationDate" orm:"expiration_date" description:"过期日期"`        // 过期日期
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`          // 备注
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`        // 创建时间
	CreatedBy      uint        `json:"createdBy"      orm:"created_by"      description:"创建人"`         // 创建人
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"修改时间"`        // 修改时间
	DeletedAt      uint        `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`        // 删除时间
}

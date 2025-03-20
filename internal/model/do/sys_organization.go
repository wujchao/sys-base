// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure of table sys_organization for DAO operations like Where/Data.
type SysOrganization struct {
	g.Meta         `orm:"table:sys_organization, do:true"`
	Id             interface{} // ID
	ParentId       interface{} // 父id
	Path           interface{} // 级别路径
	Name           interface{} // 名称
	Number         interface{} // 编号
	ListOrder      interface{} // 排序
	Leader         interface{} // 负责人
	Phone          interface{} // 联系电话
	Email          interface{} // 邮箱
	Status         interface{} // 状态（0停用 1正常）
	EffectiveDate  *gtime.Time // 生效日期
	ExpirationDate *gtime.Time // 过期日期
	Remark         interface{} // 备注
	CreatedAt      *gtime.Time // 创建时间
	CreatedBy      interface{} // 创建人
	UpdatedAt      *gtime.Time // 修改时间
	DeletedAt      interface{} // 删除时间
}

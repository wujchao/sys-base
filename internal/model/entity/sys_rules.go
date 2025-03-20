// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRules is the golang structure for table sys_rules.
type SysRules struct {
	Id         string `json:"id"         orm:"id"          description:"权限标识 全局唯一"`           // 权限标识 全局唯一
	ParentId   string `json:"parentId"   orm:"parent_id"   description:"父ID"`                 // 父ID
	Name       string `json:"name"       orm:"name"        description:"名称"`                  // 名称
	Apis       string `json:"apis"       orm:"apis"        description:"api地址列表"`             // api地址列表
	Menus      string `json:"menus"      orm:"menus"       description:"权限对应显示的菜单"`           // 权限对应显示的菜单
	FrontIds   string `json:"frontIds"   orm:"front_ids"   description:"前置权限"`                // 前置权限
	ListOrder  uint   `json:"listOrder"  orm:"list_order"  description:"排序，从小到大"`             // 排序，从小到大
	ModuleType int    `json:"moduleType" orm:"module_type" description:"类型 1全部 2 管理 3 企业"`    // 类型 1全部 2 管理 3 企业
	AuthType   uint   `json:"authType"   orm:"auth_type"   description:"授权类型，1 授权，2 登录，3 公开"` // 授权类型，1 授权，2 登录，3 公开
	Status     uint   `json:"status"     orm:"status"      description:"状态"`                  // 状态
	Desc       string `json:"desc"       orm:"desc"        description:"说明"`                  // 说明
	Remark     string `json:"remark"     orm:"remark"      description:"备注"`                  // 备注
}

package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminListParams struct {
	UserName     string `json:"userName"       orm:"user_name"       description:"用户名"`                          // 用户名
	Mobile       string `json:"mobile"         orm:"mobile"          description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname string `json:"userNickname"   orm:"user_nickname"   description:"用户昵称"`                         // 用户昵称
	UserEmail    string `json:"userEmail"      orm:"user_email"      description:"用户登录邮箱"`                       // 用户登录邮箱
	Sex          int    `json:"sex"            orm:"sex"             description:"性别0保密1男2女"`                    // 性别0保密1男2女
	IsAdmin      int    `json:"isAdmin"        orm:"is_admin"        description:"是否后台管理员 1 是  2 否"`             // 是否后台管理员 1 是  2 否
	Status       uint   `json:"status"         orm:"status"          description:"用户状态; 1:正常,2:禁用,3:未验证"`        // 用户状态; 1:正常,2:禁用,3:未验证
	*PaginationInput
}

type AdminOutput struct {
	Id             string      `json:"id"             orm:"id"              description:"ID"`                           // ID
	UserName       string      `json:"userName"       orm:"user_name"       description:"用户名"`                          // 用户名
	UserTypes      int         `json:"userTypes"      orm:"user_types"      description:"1 管理，2 企业，9系统"`                // 1 管理，2 企业，9系统
	Mobile         string      `json:"mobile"         orm:"mobile"          description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname   string      `json:"userNickname"   orm:"user_nickname"   description:"用户昵称"`                         // 用户昵称
	Sex            int         `json:"sex"            orm:"sex"             description:"性别0保密1男2女"`                    // 性别0保密1男2女
	UserEmail      string      `json:"userEmail"      orm:"user_email"      description:"用户登录邮箱"`                       // 用户登录邮箱
	Avatar         string      `json:"avatar"         orm:"avatar"          description:"用户头像"`                         // 用户头像
	DeptId         string      `json:"deptId"         orm:"dept_id"         description:"部门id"`                         // 部门id
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`                           // 备注
	IsAdmin        int         `json:"isAdmin"        orm:"is_admin"        description:"是否后台管理员 1 是  2 否"`             // 是否后台管理员 1 是  2 否
	Address        string      `json:"address"        orm:"address"         description:"联系地址"`                         // 联系地址
	LastLoginIp    string      `json:"lastLoginIp"    orm:"last_login_ip"   description:"最后登录ip"`                       // 最后登录ip
	LastLoginTime  *gtime.Time `json:"lastLoginTime"  orm:"last_login_time" description:"最后登录时间"`                       // 最后登录时间
	Status         uint        `json:"status"         orm:"status"          description:"用户状态; 1:正常,2:禁用,3:未验证"`        // 用户状态; 1:正常,2:禁用,3:未验证
	EffectiveDate  *gtime.Time `json:"effectiveDate"  orm:"effective_date"  description:"生效日期"`                         // 生效日期
	ExpirationDate *gtime.Time `json:"expirationDate" orm:"expiration_date" description:"过期日期"`                         // 过期日期
	CreatedBy      string      `json:"createdBy"      orm:"created_by"      description:"创建者"`                          // 创建者
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建日期"`                         // 创建日期
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"修改日期"`                         // 修改日期
}

type AdminListOutput struct {
	List []*AdminOutput
	PaginationOutput
}

type AdminInput struct {
	UserName       string      `json:"userName"       orm:"user_name"       description:"用户名"`                          // 用户名
	Mobile         string      `json:"mobile"         orm:"mobile"          description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname   string      `json:"userNickname"   orm:"user_nickname"   description:"用户昵称"`                         // 用户昵称
	Sex            int         `json:"sex"            orm:"sex"             description:"性别0保密1男2女"`                    // 性别0保密1男2女
	UserEmail      string      `json:"userEmail"      orm:"user_email"      description:"用户登录邮箱"`                       // 用户登录邮箱
	Avatar         string      `json:"avatar"         orm:"avatar"          description:"用户头像"`                         // 用户头像
	DeptId         string      `json:"deptId"         orm:"dept_id"         description:"部门id"`                         // 部门id
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`                           // 备注
	IsAdmin        int         `json:"isAdmin"        orm:"is_admin"        description:"是否后台管理员 1 是  2 否"`             // 是否后台管理员 1 是  2 否
	Address        string      `json:"address"        orm:"address"         description:"联系地址"`                         // 联系地址
	Status         uint        `json:"status"         orm:"status"          description:"用户状态; 1:正常,2:禁用,3:未验证"`        // 用户状态; 1:正常,2:禁用,3:未验证
	EffectiveDate  *gtime.Time `json:"effectiveDate"  orm:"effective_date"  description:"生效日期"`                         // 生效日期
	ExpirationDate *gtime.Time `json:"expirationDate" orm:"expiration_date" description:"过期日期"`                         // 过期日期
	RoleIds        []string    `json:"roleIds" description:"角色ID数组" v:"required-without:IsAdmin#角色不能为空"`
}

type AdminEditInput struct {
	Id string `json:"id"`
	*AdminInput
}

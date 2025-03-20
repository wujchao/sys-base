package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginUserRes struct {
	Id           string      `json:"id"             orm:"id"              description:"ID"`                                                                                                                                         // ID
	OrgId        string      `json:"orgId"          orm:"org_id"          description:"Enterprise ID"`                                                                                                                              // Enterprise ID
	UserName     string      `json:"userName"       orm:"user_name"       description:"Username"`                                                                                                                                   // Username
	UserTypes    int         `json:"userTypes"      orm:"user_types"      description:"1: Admin, 2: Enterprise, 9: System"`                                                                                                         // 1: Admin, 2: Enterprise, 9: System
	Mobile       string      `json:"mobile"         orm:"mobile"          description:"For Chinese mobile numbers, do not include the country code. For international mobile numbers, the format is: Country code - Mobile number"` // For Chinese mobile numbers, do not include the country code. For international mobile numbers, the format is: Country code - Mobile number
	UserNickname string      `json:"userNickname"   orm:"user_nickname"   description:"User nickname"`                                                                                                                              // User nickname
	Sex          int         `json:"sex"            orm:"sex"             description:"0: Secret, 1: Male, 2: Female"`                                                                                                              // 0: Secret, 1: Male, 2: Female
	UserEmail    string      `json:"userEmail"      orm:"user_email"      description:"User login email"`                                                                                                                           // User login email
	Avatar       string      `json:"avatar"         orm:"avatar"          description:"User avatar"`                                                                                                                                // User avatar
	DeptId       string      `json:"deptId"         orm:"dept_id"         description:"Department ID"`                                                                                                                              // Department ID
	Remark       string      `json:"remark"         orm:"remark"          description:"Remarks"`                                                                                                                                    // Remarks
	IsAdmin      int         `json:"isAdmin"        orm:"is_admin"        description:"Whether it is a background administrator. 1: Yes, 2: No"`                                                                                    // Whether it is a background administrator. 1: Yes, 2: No
	CreatedAt    *gtime.Time `json:"createdAt"      orm:"created_at"      description:"Creation date"`                                                                                                                              // Creation date
	UpdatedAt    *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"Modification date"`                                                                                                                          // Modification date
}

type UserListParams struct {
	OrgIds   string `json:"orgIds" description:"OrgIds,逗号分隔" type:"number[]|string"`
	Keywords string `json:"keywords" description:"关键词(可根据账号或者用户昵称查询)"`
	UserName string `json:"userName"  description:"用户名"`
	Mobile   string `json:"mobile"  description:"手机号"`
	Status   int    `json:"status"  description:"用户状态; 1:正常,2:禁用,3:未验证"`
	*PaginationInput
}

type UserOutput struct {
	g.Meta        `orm:"table:sys_users, do:true"`
	Id            string           `json:"id"            orm:"id"              ` // ID
	OrgId         string           `json:"orgId"         orm:"org_id"          ` // 企业ID
	UserName      string           `json:"userName"      orm:"user_name"       ` // 用户名
	UserTypes     int              `json:"userTypes"     orm:"user_types"      ` // 1 系统，2 企业
	Mobile        string           `json:"mobile"        orm:"mobile"          ` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string           `json:"userNickname"  orm:"user_nickname"   ` // 用户昵称
	UserEmail     string           `json:"userEmail"     orm:"user_email"      ` // 用户登录邮箱
	Sex           int              `json:"sex"           orm:"sex"             ` // 性别;0:保密,1:男,2:女
	Avatar        string           `json:"avatar"        orm:"avatar"          ` // 用户头像
	DeptId        string           `json:"deptId"        orm:"dept_id"         ` // 部门id
	Remark        string           `json:"remark"        orm:"remark"          ` // 备注
	IsAdmin       int              `json:"isAdmin"       orm:"is_admin"        ` // 是否后台管理员 1 是  2 否
	Address       string           `json:"address"       orm:"address"         ` // 联系地址
	LastLoginIp   string           `json:"lastLoginIp"   orm:"last_login_ip"   ` // 最后登录ip
	LastLoginTime *gtime.Time      `json:"lastLoginTime" orm:"last_login_time" ` // 最后登录时间
	Status        uint             `json:"status"        orm:"status"          ` // 用户状态; 1:正常,2:禁用,3:未验证
	CreatedBy     uint64           `json:"createdBy"     orm:"created_by"      ` // 创建者
	CreatedAt     *gtime.Time      `json:"createdAt"     orm:"created_at"      ` // 创建日期
	UpdatedAt     *gtime.Time      `json:"updatedAt"     orm:"updated_at"      ` // 修改日期
	Roles         []*SysRoleOutput `json:"roles" dc:"角色" orm:"with:id=sys_users_role.user_id;sys_users_role.role_id=roles.id"`
}

type UserListOutput struct {
	List []*UserOutput `json:"list" dc:"用户列表"`
	PaginationOutput
}

type UserInput struct {
	OrgId        string   `json:"orgId"         description:"OrgId" v:"required#组织不能为空"          `
	UserName     string   `json:"userName"      description:"用户名" v:"required#用户名不能为空"`
	UserTypes    *int     `json:"userTypes"     description:"1 管理，2 企业，9系统" d:"2"`
	Mobile       string   `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string   `json:"userNickname"  description:"用户昵称"`
	UserPassword string   `json:"userPassword"  description:"登录密码" v:"required#密码不能为空"`
	UserEmail    string   `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int      `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string   `json:"avatar"        description:"用户头像"`
	DeptId       string   `json:"deptId"        description:"部门id"`
	Remark       string   `json:"remark"        description:"备注"`
	IsAdmin      int      `json:"isAdmin"       description:"是否后台管理员 1 是  2 否，只有管理员才能创建管理员" d:"2"`
	Address      string   `json:"address"       description:"联系地址"`
	Status       uint     `json:"status"        description:"用户状态; 1:正常,2:禁用,3:未验证" d:"1"`
	RoleIds      []string `json:"roleIds"       description:"角色ID数组" v:"required-without:IsAdmin#角色不能为空"`
}

type UserEditInput struct {
	Id           string   `json:"id"            description:"Id" v:"required#ID不能为空"`
	OrgId        string   `json:"orgId"         description:"OrgId" v:"required#orgId不能为空"          ` // 企业ID
	UserName     string   `json:"userName"      description:"用户名"`
	Mobile       string   `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string   `json:"userNickname"  description:"用户昵称"`
	UserPassword string   `json:"userPassword"  description:"登录密码"`
	UserEmail    string   `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int      `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string   `json:"avatar"        description:"用户头像"`
	DeptId       string   `json:"deptId"        description:"部门id"`
	Remark       string   `json:"remark"        description:"备注"`
	IsAdmin      int      `json:"isAdmin"       description:"是否后台管理员 1 是  2 否，只有管理员才能创建管理员"`
	Address      string   `json:"address"       description:"联系地址"`
	Status       uint     `json:"status"        description:"用户状态; 1:正常,2:禁用,3:未验证"`
	RoleIds      []string `json:"roleIds"      description:"角色ID数组" v:"required-without:IsAdmin#角色不能为空"`
}

type UserSelfOutput struct {
	Id            string           `json:"id"            orm:"id"              ` // ID
	OrgId         string           `json:"orgId"         orm:"org_id"          ` // 企业ID
	UserName      string           `json:"userName"      orm:"user_name"       ` // 用户名
	UserTypes     *int             `json:"userTypes"     orm:"user_types"      ` // 1 系统，2 企业
	Mobile        string           `json:"mobile"        orm:"mobile"          ` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string           `json:"userNickname"  orm:"user_nickname"   ` // 用户昵称
	UserEmail     string           `json:"userEmail"     orm:"user_email"      ` // 用户登录邮箱
	Sex           int              `json:"sex"           orm:"sex"             ` // 性别;0:保密,1:男,2:女
	Avatar        string           `json:"avatar"        orm:"avatar"          ` // 用户头像
	DeptId        string           `json:"deptId"        orm:"dept_id"         ` // 部门id
	Remark        string           `json:"remark"        orm:"remark"          ` // 备注
	IsAdmin       int              `json:"isAdmin"       orm:"is_admin"        ` // 是否后台管理员 1 是  2 否
	Address       string           `json:"address"       orm:"address"         ` // 联系地址
	LastLoginIp   string           `json:"lastLoginIp"   orm:"last_login_ip"   ` // 最后登录ip
	LastLoginTime *gtime.Time      `json:"lastLoginTime" orm:"last_login_time" ` // 最后登录时间
	Status        uint             `json:"status"        orm:"status"          ` // 用户状态; 1:正常,2:禁用,3:未验证
	CreatedBy     string           `json:"createdBy"     orm:"created_by"      ` // 创建者
	CreatedAt     *gtime.Time      `json:"createdAt"     orm:"created_at"      ` // 创建日期
	UpdatedAt     *gtime.Time      `json:"updatedAt"     orm:"updated_at"      ` // 修改日期
	Roles         []*SysRoleOutput `json:"roles"      description:"角色数组" v:"required#角色不能为空"`
}

type UserSelfEditInput struct {
	Mobile       string `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string `json:"userNickname"  description:"用户昵称"`
	UserEmail    string `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int    `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string `json:"avatar"        description:"用户头像"`
	Address      string `json:"address"       description:"联系地址"`
	Remark       string `json:"remark"        description:"备注"`
}

type UserSelfChanPasswordInput struct {
	Password    string `json:"password" v:"required#old password cannot be empty"`
	OldPassword string `json:"oldPassword" v:"required#new password cannot be empty"`
}

type UserUpdateLoginInput struct {
	Id            string      `json:"id"`
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   ` // 最后登录ip
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" ` // 最后登录时间
}

type CheckPermissionsInput struct {
	OrgId     string `json:"orgId"`
	UserTypes *int   `json:"userTypes,omitempty"          ` // 1 系统，2 企业
	IsAdmin   int    `json:"isAdmin,omitempty"        `     // 是否后台管理员 1 是  2 否
}

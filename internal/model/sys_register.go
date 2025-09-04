package model

type RegisterInput struct {
	UserName     string `json:"userName" v:"required#用户名不能为空" dc:"用户名"`
	UserPassword string `json:"userPassword" v:"required#密码不能为空" dc:"密码"`
	Mobile       string `json:"mobile,omitempty" dc:"手机号码"`
	OrgId        string `json:"orgId,omitempty" dc:"企业ID"`
	UserNickname string `json:"userNickname,omitempty" dc:"用户昵称"`
	UserEmail    string `json:"userEmail,omitempty" dc:"用户邮箱"`
	Sex          int    `json:"sex,omitempty" dc:"性别"`
	UserTypes    *int   `json:"userTypes,omitempty" d:"3" dc:"3 用户, 只支持用户注册"`
	Avatar       string `json:"avatar,omitempty" dc:"头像"`
	DeptId       string `json:"deptId,omitempty" dc:"部门ID"`
	IsAdmin      int    `json:"isAdmin,omitempty" d:"2" dc:"是否是管理员"`
	Remark       string `json:"remark,omitempty" dc:"备注"`
	Address      string `json:"address,omitempty" dc:"地址"`
}

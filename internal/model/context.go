package model

import "github.com/gogf/gf/v2/frame/g"

type Context struct {
	User    *ContextUser `json:"user"`
	Data    g.Map        `json:"data"`
	AuthWay string       `json:"authWay"`
}

type ContextUser struct {
	Id           string `json:"id"             orm:"id"              description:"ID"`                                                                                                                                         // ID
	OrgId        string `json:"orgId"          orm:"org_id"          description:"Enterprise ID"`                                                                                                                              // Enterprise ID
	UserName     string `json:"userName"       orm:"user_name"       description:"Username"`                                                                                                                                   // Username
	UserTypes    int    `json:"userTypes"      orm:"user_types"      description:"1: Admin, 2: Enterprise, 9: System"`                                                                                                         // 1: Admin, 2: Enterprise, 9: System
	Mobile       string `json:"mobile"         orm:"mobile"          description:"For Chinese mobile numbers, do not include the country code. For international mobile numbers, the format is: Country code - Mobile number"` // For Chinese mobile numbers, do not include the country code. For international mobile numbers, the format is: Country code - Mobile number
	UserNickname string `json:"userNickname"   orm:"user_nickname"   description:"User nickname"`                                                                                                                              // User nickname
	Sex          int    `json:"sex"            orm:"sex"             description:"0: Secret, 1: Male, 2: Female"`                                                                                                              // 0: Secret, 1: Male, 2: Female
	UserEmail    string `json:"userEmail"      orm:"user_email"      description:"User login email"`                                                                                                                           // User login email
	Avatar       string `json:"avatar"         orm:"avatar"          description:"User avatar"`                                                                                                                                // User avatar
	DeptId       string `json:"deptId"         orm:"dept_id"         description:"Department ID"`                                                                                                                              // Department ID
	Remark       string `json:"remark"         orm:"remark"          description:"Remarks"`                                                                                                                                    // Remarks
	IsAdmin      int    `json:"isAdmin"        orm:"is_admin"        description:"Whether it is a background administrator. 1: Yes, 2: No"`                                                                                    // Whether it is a background administrator. 1: Yes, 2: No
}

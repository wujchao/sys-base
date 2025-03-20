package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/internal/model"
)

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" rules:"UserRead" parent:"User" menu:"System OrgUser User" summary:"User List" tags:"User Management"`
	*model.UserListParams
}
type UserListRes struct {
	*model.UserListOutput
}

type UserReadReq struct {
	g.Meta `path:"/user/read" method:"get" rules:"UserRead" parent:"User" menu:"System OrgUser User" summary:"User Read" tags:"User Management"`
	Id     string `json:"id" v:"required#ID cannot be empty"`
}

type UserReadRes model.UserOutput

type UserAddReq struct {
	g.Meta `path:"/user/add" method:"post" rules:"UserAdd" parent:"User" menu:"System OrgUser User" summary:"User Add" tags:"User Management"`
	*model.UserInput
}
type UserAddRes struct {
	Id string `json:"id"`
}

type UserEditReq struct {
	g.Meta `path:"/user/edit" method:"put" rules:"UserEdit" parent:"User" menu:"System OrgUser User"  summary:"User Edit" tags:"User Management"`
	*model.UserEditInput
}
type UserEditRes struct {
	Id string `json:"id"`
}

type UserDeleteByIdReq struct {
	g.Meta `path:"/user/delete" method:"delete" rules:"UserDel" parent:"User" menu:"System OrgUser User" summary:"Delete user by ID" tags:"User Management"`
	Id     string `json:"id"        description:"ID" v:"required#ID cannot be empty"`
}
type UserDeleteByIdRes struct {
	Id string `json:"id"`
}

type ResetPasswordReq struct {
	g.Meta      `path:"/user/resetPassword" method:"post" rules:"UserResetPass" parent:"User" menu:"System OrgUser User" summary:"Reset User Password" tags:"User Management"`
	Id          string `json:"id"        description:"ID" v:"required#ID cannot be empty"`
	NewPassword string `json:"newPassword"  description:"reset login password;"`
}

type ResetPasswordRes struct {
}

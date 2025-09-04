package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wujchao/sys-base/internal/model"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" rules:"RolesRead" parent:"Roles" menu:"System SystemRole" summary:"Role list" tags:"Role Management"`
	*model.RoleListReqParam
}

type RoleListRes struct {
	*model.SysRoleListOutput
}

type RoleReadReq struct {
	g.Meta `path:"/role/read" method:"get" rules:"RolesRead" parent:"Roles" menu:"System SystemRole" summary:"Role query" tags:"Role Management"`
	Id     string `json:"id"`
}

type RoleReadRes model.SysRoleOutput

type RoleAddReq struct {
	g.Meta `path:"/role/add" method:"post" rules:"RolesAdd" parent:"Roles" menu:"System SystemRole" front:"RolesRead" summary:"Add role" tags:"Role Management"`
	*model.SysRoleInput
}
type RoleAddRes struct {
	Id string `json:"id"`
}

type RoleEditReq struct {
	g.Meta `path:"/role/edit" method:"put" rules:"RolesEdit" parent:"Roles" menu:"System SystemRole" front:"RolesRead" summary:"Edit Role" tags:"Role Management"`
	*model.SysRoleEditInput
}

type RoleEditRes struct {
	Id string `json:"id"`
}

type RoleDeleteByIdReq struct {
	g.Meta `path:"/role/delete" method:"delete" rules:"RolesDel" parent:"Roles" menu:"System SystemRole" front:"RolesRead" summary:"Delete Role" tags:"Role Management"`
	Id     string `json:"id"`
}

type RoleDeleteByIdRes struct {
	Id string `json:"id"`
}

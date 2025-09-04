package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wujchao/sys-base/internal/model"
)

type AdminListReq struct {
	g.Meta `path:"/admin/list" method:"get" rules:"AdminRead" parent:"Admin" menu:"System Admin" module:"admin" summary:"Admin List" tags:"Admin Management"`
	*model.AdminListParams
}
type AdminListRes struct {
	*model.AdminListOutput
}

type AdminReadReq struct {
	g.Meta `path:"/admin/read" method:"get" rules:"AdminRead" parent:"Admin" menu:"System Admin" module:"admin" summary:"Admin Read" tags:"Admin Management"`
	Id     string `json:"id"`
}

type AdminReadRes model.AdminOutput

type AdminAddReq struct {
	g.Meta `path:"/admin/add" method:"post" rules:"AdminAdd" parent:"Admin" menu:"System Admin" module:"admin" summary:"Admin Add" tags:"Admin Management"`
	*model.AdminInput
}
type AdminAddRes struct {
	Id string `json:"id"`
}

type AdminEditReq struct {
	g.Meta `path:"/admin/edit" method:"put" rules:"AdminEdit" parent:"Admin" menu:"System Admin" module:"admin" summary:"Admin Edit" tags:"Admin Management"`
	*model.AdminEditInput
}
type AdminEditRes struct {
	Id string `json:"id"`
}

type AdminDeleteByIdReq struct {
	g.Meta `path:"/admin/delete" method:"delete" rules:"AdminDel" parent:"Admin" menu:"System Admin" module:"admin" summary:"Delete admin by ID" tags:"Admin Management"`
	Id     string `json:"id"        description:"ID" v:"required#ID cannot be empty"`
}
type AdminDeleteByIdRes struct {
	Id string `json:"id"`
}

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wujchao/sys-base/internal/model"
)

type OneselfMenuReq struct {
	g.Meta `path:"/oneself/menus" method:"get" rules:"SysMenuTree" auth:"login" tags:"Oneself" summary:"Oneself Menu Tree"`
}

type OneselfMenuRes []*model.SysMenuRes

type OneselfRulesTreeReq struct {
	g.Meta   `path:"/oneself/rules" method:"get" rules:"SysRulesTree" auth:"login" summary:"Oneself Rule Tree" tags:"Oneself"`
	ParentId *uint64 `json:"parentId" dc:"parent ID"` //
}

type OneselfRulesTreeRes []*model.SysRulesTreeOutput

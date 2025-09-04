package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wujchao/sys-base/internal/model"
)

type SysDocReadReq struct {
	g.Meta `path:"/sys/doc/read" method:"post" tags:"SysDoc" summary:"Read doc"`
	Key    string `json:"key"`
}

type SysDocReadRes model.SysDocOutput

package model

import "github.com/wujchao/sys-base/internal/model/entity"

type SysRulesOutput struct {
	*entity.SysRules
}

type SysRulesTreeOutput struct {
	*SysRulesOutput
	Children []*SysRulesTreeOutput `json:"children"`
}

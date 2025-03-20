package model

import "sys-base/internal/model/entity"

type SysRulesOutput struct {
	*entity.SysRules
}

type SysRulesTreeOutput struct {
	*SysRulesOutput
	Children []*SysRulesTreeOutput `json:"children"`
}

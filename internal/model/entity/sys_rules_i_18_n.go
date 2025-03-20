// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRulesI18N is the golang structure for table sys_rules_i18n.
type SysRulesI18N struct {
	Id       uint   `json:"id"       orm:"id"        description:""`     //
	LangCode string `json:"langCode" orm:"lang_code" description:"语言代码"` // 语言代码
	Name     string `json:"name"     orm:"name"      description:"名称"`   // 名称
	Desc     string `json:"desc"     orm:"desc"      description:"说明"`   // 说明
}

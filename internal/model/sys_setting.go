package model

type SettingInput struct {
	Id           string `json:"id"           orm:"id"            ` // 系统设置key
	SettingGroup string `json:"settingGroup" orm:"setting_group" ` // 设置分组
	Value        string `json:"value"        orm:"value"         ` // 设置值
	Name         string `json:"name"         orm:"name"          ` // 设置key
	Description  string `json:"description"  orm:"description"   ` // 描述
	Status       int    `json:"status"       orm:"status"        ` // 状态1启用2禁用
	SortList     int    `json:"sortList"     orm:"sort_list"     ` // 排序
}

type SettingByGroupParams struct {
	SettingGroup string `json:"settingGroup" `
}

type SettingOutput struct {
	Id           string `json:"id"           orm:"id"            ` // 系统设置key
	SettingGroup string `json:"settingGroup" orm:"setting_group" ` // 设置分组
	Value        string `json:"value"        orm:"value"         ` // 设置值
	Name         string `json:"name"         orm:"name"          ` // 设置key
	Description  string `json:"description"  orm:"description"   ` // 描述
	Type         int    `json:"type"         orm:"type"          ` // 1系统内置2否
	Status       int    `json:"status"       orm:"status"        ` // 状态1启用2禁用
	SortList     int    `json:"sortList"     orm:"sort_list"     ` // 排序
}

type SettingListOutput struct {
	List []*SettingOutput
}

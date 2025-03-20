package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/internal/model"
)

type SettingAddReq struct {
	g.Meta `path:"/setting/add" method:"post" rules:"SysSettingAdd" parent:"SysSetting" summary:"Add Settings" menu:"System SysSettings" tags:"Settings Management"`
	*model.SettingInput
}

type SettingAddRes struct {
}

type SettingGetByIdReq struct {
	g.Meta `path:"/setting/getByKey" method:"get" rules:"SysSettingRead" parent:"SysSetting" summary:"Get settings based on Key" tags:"Settings Management"`
	Id     string `json:"key" description:"Key" v:"required#Key cannot be empty"`
}

type SettingGetByIdRes struct {
	*model.SettingOutput
}

type SettingGetByGroupReq struct {
	g.Meta `path:"/setting/getByGroup" method:"get" rules:"SysSettingRead" parent:"SysSetting" summary:"Get settings by group" tags:"Settings Management"`
	Group  string `json:"group" description:"Grouping" v:"required#Group cannot be empty"`
}

type SettingGetByGroupRes struct {
	*model.SettingListOutput
}

type SettingDeleteByIdReq struct {
	g.Meta `path:"/setting/deleteById" method:"delete" rules:"SysSettingDel" parent:"SysSetting" summary:"Delete Settings" tags:"Settings Management"`
	Id     string `json:"id" description:"ID" v:"required#ID cannot be empty"`
}

type SettingDeleteByIdRes struct {
	Id string `json:"id"`
}

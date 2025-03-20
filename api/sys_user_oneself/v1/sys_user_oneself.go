package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/internal/model"
)

type SelfUserReadReq struct {
	g.Meta `path:"/selfUser/read" method:"get" rules:"SelfUserRead" auth:"login" summary:"User Read" tags:"User Oneself"`
}
type SelfUserReadRes model.UserSelfOutput

type SelfUserEditReq struct {
	g.Meta `path:"/selfUser/edit" method:"put" rules:"SelfUserEdit" auth:"login" summary:"User Edit" tags:"User Oneself"`
	*model.UserSelfEditInput
}

type SelfUserEditRes struct {
}

type SelfUserChanPassReq struct {
	g.Meta `path:"/selfUser/chanPass" method:"put" rules:"SelfUserEdit" auth:"login" summary:"User Change Password" tags:"User Oneself"`
	*model.UserSelfChanPasswordInput
}

type SelfUserChanPassRes struct {
}

package sys_user_oneself

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_user_oneself/v1"
)

func (c *ControllerV1) SelfUserChanPass(ctx context.Context, req *v1.SelfUserChanPassReq) (res *v1.SelfUserChanPassRes, err error) {
	err = service.SysUsers().ChanPassSelf(ctx, req.UserSelfChanPasswordInput)
	return
}

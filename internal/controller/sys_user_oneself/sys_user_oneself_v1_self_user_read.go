package sys_user_oneself

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/consts"
	"sys-base/service"

	"sys-base/api/sys_user_oneself/v1"
)

func (c *ControllerV1) SelfUserRead(ctx context.Context, req *v1.SelfUserReadReq) (res *v1.SelfUserReadRes, err error) {
	id := service.Context().GetUserId(ctx)
	if id == nil {
		return nil, gerror.NewCode(consts.NotLoggedInErrCode)
	}
	data, err := service.SysUsers().Read(ctx, *id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	err = gconv.Scan(data, &res)
	return
}

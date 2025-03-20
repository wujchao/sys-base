package sys_user

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_user/v1"
)

func (c *ControllerV1) UserAdd(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	id, err := service.SysUsers().Add(ctx, req.UserInput)
	if err != nil {
		return nil, err
	}
	res = &v1.UserAddRes{
		Id: id,
	}
	return
}

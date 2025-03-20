package sys_user

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_user/v1"
)

func (c *ControllerV1) UserRead(ctx context.Context, req *v1.UserReadReq) (res *v1.UserReadRes, err error) {
	data, err := service.SysUsers().Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	d := v1.UserReadRes(*data)
	res = &d
	return
}

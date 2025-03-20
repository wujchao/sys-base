package sys_login

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_login/v1"
)

func (c *ControllerV1) LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res, err = service.Login().Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

package sys_login

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_login/v1"
)

func (c *ControllerV1) LoginOut(ctx context.Context, req *v1.LoginOutReq) (res *v1.LoginOutRes, err error) {
	err = service.Login().Logout(ctx)
	return
}

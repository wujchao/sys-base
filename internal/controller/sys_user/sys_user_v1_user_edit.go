package sys_user

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_user/v1"
)

func (c *ControllerV1) UserEdit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserEditRes, err error) {
	err = service.SysUsers().Edit(ctx, req.UserEditInput)
	if err != nil {
		return nil, err
	}
	res = &v1.UserEditRes{Id: req.UserEditInput.Id}
	return
}

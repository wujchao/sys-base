package sys_user

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_user/v1"
)

func (c *ControllerV1) UserDeleteById(ctx context.Context, req *v1.UserDeleteByIdReq) (res *v1.UserDeleteByIdRes, err error) {
	err = service.SysUsers().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.UserDeleteByIdRes{Id: req.Id}
	return
}

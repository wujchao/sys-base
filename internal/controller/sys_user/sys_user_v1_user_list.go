package sys_user

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_user/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	data, err := service.SysUsers().ListOneself(ctx, req.UserListParams)
	if err != nil {
		return nil, err
	}
	res = &v1.UserListRes{UserListOutput: data}
	return
}

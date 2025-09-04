package sys_role

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_role/v1"
)

func (c *ControllerV1) RoleAdd(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error) {
	data, err := service.SysRole().Add(ctx, req.SysRoleInput)
	if err != nil {
		return nil, err
	}
	res = &v1.RoleAddRes{
		Id: data.Id,
	}
	return
}

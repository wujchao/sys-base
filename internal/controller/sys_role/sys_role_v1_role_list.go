package sys_role

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_role/v1"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	data, err := service.SysRole().List(ctx, req.RoleListReqParam)
	if err != nil {
		return nil, err
	}
	res = &v1.RoleListRes{
		SysRoleListOutput: data,
	}
	return
}

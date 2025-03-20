package sys_role

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_role/v1"
)

func (c *ControllerV1) RoleEdit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleEditRes, err error) {
	err = service.SysRole().Edit(ctx, req.SysRoleEditInput)
	if err != nil {
		return nil, err
	}
	res = &v1.RoleEditRes{Id: req.Id}
	return
}

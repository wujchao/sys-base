package sys_role

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_role/v1"
)

func (c *ControllerV1) RoleDeleteById(ctx context.Context, req *v1.RoleDeleteByIdReq) (res *v1.RoleDeleteByIdRes, err error) {
	err = service.SysRole().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.RoleDeleteByIdRes{Id: req.Id}
	return
}

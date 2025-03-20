package sys_admin

import (
	"context"
	"sys-base/service"

	"sys-base/api/sys_admin/v1"
)

func (c *ControllerV1) AdminDeleteById(ctx context.Context, req *v1.AdminDeleteByIdReq) (res *v1.AdminDeleteByIdRes, err error) {
	err = service.SysUsers().Delete(ctx, req.Id)
	return
}

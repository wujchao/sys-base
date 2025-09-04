package sys_role

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_role/v1"
)

func (c *ControllerV1) RoleRead(ctx context.Context, req *v1.RoleReadReq) (res *v1.RoleReadRes, err error) {
	data, err := service.SysRole().Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &res)
	return
}

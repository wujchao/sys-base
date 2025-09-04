package sys_admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_admin/v1"
)

func (c *ControllerV1) AdminRead(ctx context.Context, req *v1.AdminReadReq) (res *v1.AdminReadRes, err error) {
	data, err := service.SysUsers().Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &res)
	return
}

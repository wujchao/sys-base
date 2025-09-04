package sys_doc

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/service"

	"sys-base/api/sys_doc/v1"
)

func (c *ControllerV1) SysDocRead(ctx context.Context, req *v1.SysDocReadReq) (res *v1.SysDocReadRes, err error) {
	data, err := service.SysDoc().Read(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &res)
	return
}

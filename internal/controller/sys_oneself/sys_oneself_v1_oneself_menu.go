package sys_oneself

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_oneself/v1"
)

func (c *ControllerV1) OneselfMenu(ctx context.Context, req *v1.OneselfMenuReq) (res *v1.OneselfMenuRes, err error) {
	// 获取菜单
	data, err := service.SysMenu().GetTree(ctx)
	if err != nil {
		return nil, err
	}
	d := v1.OneselfMenuRes(data)
	return &d, nil
}

package sys_oneself

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_oneself/v1"
)

func (c *ControllerV1) OneselfRulesTree(ctx context.Context, req *v1.OneselfRulesTreeReq) (res *v1.OneselfRulesTreeRes, err error) {
	data, err := service.SysRules().Tree(ctx)
	if err != nil {
		return nil, err
	}
	d := v1.OneselfRulesTreeRes(data)
	return &d, nil
}

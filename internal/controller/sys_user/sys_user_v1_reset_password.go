package sys_user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/wujchao/sys-base/api/sys_user/v1"
)

func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

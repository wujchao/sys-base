package sys_setting

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sys-base/api/sys_setting/v1"
)

func (c *ControllerV1) SettingAdd(ctx context.Context, req *v1.SettingAddReq) (res *v1.SettingAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

package sys_setting

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/wujchao/sys-base/api/sys_setting/v1"
)

func (c *ControllerV1) SettingGetById(ctx context.Context, req *v1.SettingGetByIdReq) (res *v1.SettingGetByIdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

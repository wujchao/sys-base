package sys_admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/wujchao/sys-base/api/sys_admin/v1"
)

func (c *ControllerV1) AdminList(ctx context.Context, req *v1.AdminListReq) (res *v1.AdminListRes, err error) {
	userListParams := &model.UserListParams{}
	if err = gconv.Scan(req.AdminListParams, &userListParams); err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "参数错误")
	}
	data, err := service.SysUsers().ListOneself(ctx, userListParams)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &res)
	return
}

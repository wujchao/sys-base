package sys_admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/internal/model"
	"sys-base/service"

	"sys-base/api/sys_admin/v1"
)

func (c *ControllerV1) AdminEdit(ctx context.Context, req *v1.AdminEditReq) (res *v1.AdminEditRes, err error) {
	userInput := &model.UserEditInput{}
	err = gconv.Scan(req.AdminEditInput, &userInput)
	if err != nil {
		return nil, err
	}
	err = service.SysUsers().Edit(ctx, userInput)
	return
}

package sys_admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/consts"
	"sys-base/internal/model"
	"sys-base/service"

	"sys-base/api/sys_admin/v1"
)

func (c *ControllerV1) AdminAdd(ctx context.Context, req *v1.AdminAddReq) (res *v1.AdminAddRes, err error) {
	userType := consts.UserTypeAdmin
	userInput := &model.UserInput{
		OrgId:     "",
		UserTypes: &userType,
	}
	err = gconv.Scan(req.AdminInput, &userInput)
	id, err := service.SysUsers().Add(ctx, userInput)
	if err != nil {
		return nil, err
	}
	res = &v1.AdminAddRes{
		Id: id,
	}
	return
}

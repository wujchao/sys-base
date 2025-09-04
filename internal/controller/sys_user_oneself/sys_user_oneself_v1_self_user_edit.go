package sys_user_oneself

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/consts"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_user_oneself/v1"
)

func (c *ControllerV1) SelfUserEdit(ctx context.Context, req *v1.SelfUserEditReq) (res *v1.SelfUserEditRes, err error) {
	uid := service.Context().GetUserId(ctx)
	if uid == nil {
		return nil, gerror.NewCode(consts.NotLoggedInErrCode)
	}
	in := &model.UserEditInput{
		Id: *uid,
	}
	err = gconv.Struct(req.UserSelfEditInput, &in)
	if err != nil {
		return nil, err
	}
	err = service.SysUsers().Edit(ctx, in)
	if err != nil {
		return nil, err
	}
	return
}

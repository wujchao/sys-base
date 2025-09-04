package sys_upload

import (
	"context"
	"github.com/wujchao/sys-base/consts"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/wujchao/sys-base/api/sys_upload/v1"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(consts.ParamsInvalidErrCode)
	}
	result, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	return
}

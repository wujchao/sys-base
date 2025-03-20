package sys_upload

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sys-base/consts"
	"sys-base/internal/model"
	"sys-base/service"

	"sys-base/api/sys_upload/v1"
)

func (c *ControllerV1) FileUploadForEditor(ctx context.Context, req *v1.FileUploadForEditorReq) (res *v1.FileUploadForEditorRes, err error) {
	// Editor 上传图片需要规定的返回格式，在控制器直接返回
	err = service.Context().SetData(ctx, consts.DirectKey, true)
	if err != nil {
		return
	}
	data := &v1.FileUploadForEditorRes{
		Errno:   1,
		Message: g.I18n().T(ctx, "上传失败"),
	}
	if req.File == nil {
		data.Message = g.I18n().T(ctx, "上传文件不能为空")
		return data, nil
	}
	result, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return data, nil
	}
	res = &v1.FileUploadForEditorRes{
		Data: &v1.FileUploadForEditorData{
			Url:  result.Url,
			Alt:  result.Name,
			Href: result.Url,
		},
		Errno: 0,
	}
	return
}

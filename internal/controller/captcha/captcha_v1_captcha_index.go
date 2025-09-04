package captcha

import (
	"context"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/captcha/v1"
)

func (c *ControllerV1) CaptchaIndex(ctx context.Context, req *v1.CaptchaIndexReq) (res *v1.CaptchaIndexRes, err error) {
	idKey, base64String, err := service.Captcha().GetVerifyImgString(ctx)
	res = &v1.CaptchaIndexRes{
		Key: idKey,
		Img: base64String,
	}
	return
}

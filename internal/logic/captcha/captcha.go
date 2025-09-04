package captcha

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
	"github.com/wujchao/sys-base/service"
	"image/color"
	"time"
)

type sCaptcha struct {
}

func init() {
	service.RegisterCaptcha(&sCaptcha{})
}

var (
	captchaStore  = base64Captcha.NewMemoryStore(100, 60*time.Second)
	captchaDriver = &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      0, //文本噪声计数
		ShowLineOptions: 2,
		Length:          4,
		Source:          "123456789",
		Fonts:           []string{"chromohv.ttf"},
		BgColor:         &color.RGBA{R: 209, G: 205, B: 205, A: 20},
	}
)

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKey, base64String string, err error) {
	driver := captchaDriver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, captchaStore)
	idKey, base64String, _, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}

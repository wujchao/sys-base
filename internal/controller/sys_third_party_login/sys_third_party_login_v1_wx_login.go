package sys_third_party_login

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/consts"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"

	"github.com/wujchao/sys-base/api/sys_third_party_login/v1"
)

func (c *ControllerV1) WxLogin(ctx context.Context, req *v1.WxLoginReq) (res *v1.WxLoginRes, err error) {
	appId := g.Cfg().MustGet(ctx, "wx.appId").String()
	appSecret := g.Cfg().MustGet(ctx, "wx.appSecret").String()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, req.Code)
	resp, err := g.Client().Get(ctx, url)
	if err != nil {
		return
	}
	defer resp.Close()
	var wxResp model.WechatLoginResponse
	if err := gconv.Scan(resp.ReadAllString(), &wxResp); err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		return nil, gerror.NewCode(consts.SystemErrCode, wxResp.ErrMsg)
	}

	data, err := service.ThirdPartyLogin().WxLogin(ctx, &wxResp)
	if err != nil {
		return nil, err
	}
	res = &v1.WxLoginRes{
		ThirdPartyLoginOutput: data.ThirdPartyLoginOutput,
	}

	return
}

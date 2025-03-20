package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"sys-base/internal/controller/captcha"
	"sys-base/internal/controller/sys_login"
	"sys-base/internal/controller/sys_upload"
)

func publicRoute(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(
			captcha.NewV1(),
			sys_login.NewV1(),
			sys_upload.NewV1(),
		)
	})
}

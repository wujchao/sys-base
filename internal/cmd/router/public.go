package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wujchao/sys-base/internal/controller/captcha"
	"github.com/wujchao/sys-base/internal/controller/sys_doc"
	"github.com/wujchao/sys-base/internal/controller/sys_login"
	"github.com/wujchao/sys-base/internal/controller/sys_third_party_login"
	"github.com/wujchao/sys-base/internal/controller/sys_upload"
)

func publicRoute(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(
			captcha.NewV1(),
			sys_login.NewV1(),
			sys_upload.NewV1(),
			sys_third_party_login.NewV1(),
			sys_doc.NewV1(),
		)
	})
}

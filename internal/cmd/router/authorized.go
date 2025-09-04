package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wujchao/sys-base/internal/controller/sys_admin"
	"github.com/wujchao/sys-base/internal/controller/sys_role"
	"github.com/wujchao/sys-base/internal/controller/sys_setting"
	"github.com/wujchao/sys-base/internal/controller/sys_user"
)

func authRoute(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(
			sys_user.NewV1(),
			sys_admin.NewV1(),
			sys_role.NewV1(),
			sys_setting.NewV1(),
		)
	})
}

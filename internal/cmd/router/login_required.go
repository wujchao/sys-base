package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wujchao/sys-base/internal/controller/sys_oneself"
	"github.com/wujchao/sys-base/internal/controller/sys_user_oneself"
)

func loginRoute(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(
			sys_oneself.NewV1(),
			sys_user_oneself.NewV1(),
		)
	})
}

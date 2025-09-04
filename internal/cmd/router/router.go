package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wujchao/sys-base/service"
)

func LoadRouter(ctx context.Context, group *ghttp.RouterGroup) {
	group.Middleware(service.Middleware().Ctx)

	group.Group("/v1", func(g *ghttp.RouterGroup) {
		loadV1(ctx, g)
	})
}

func loadV1(ctx context.Context, group *ghttp.RouterGroup) {

	// 不需要登录
	group.Group("/", func(g *ghttp.RouterGroup) {
		publicRoute(ctx, g)
	})

	// 仅登录
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware().Auth)

		loginRoute(ctx, g)
	})

	// 鉴权路由
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware().Auth)
		g.Middleware(service.Middleware().Permission)

		authRoute(ctx, g)
	})
}

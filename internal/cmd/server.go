package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"
	"os"
	"path"
	"sys-base/internal/cmd/router"
	"sys-base/internal/logic/middleware"
	"sys-base/service"
	"syscall"
)

func RunServer(ctx context.Context, stopSignal chan os.Signal) *ghttp.Server {

	// 性能分析
	enablePProf := g.Cfg().MustGet(context.Background(), "pprofEnabled").Bool()
	if enablePProf {
		RunSysPProf(stopSignal)
	}

	s := g.Server()
	// 静态目录，vue项目
	if gfile.Exists("/resource/dist/dist-pro/") {
		s.SetServerRoot("/resource/dist/dist-pro/")
	}

	// 统一返回
	s.Use(middleware.HandlerResponse)
	// 直接返回
	s.Use(service.Middleware().DirectResponse)
	// 跨域
	s.Use(service.Middleware().MiddlewareCORS)
	// 操作日志
	s.Use(service.Middleware().OperationLog)
	// 关闭打印路由
	//s.SetDumpRouterMap(false)
	// 自定义Swig文档
	if g.Cfg().MustGet(ctx, "swagger.custom").Bool() {
		customOpenApiDoc(s)
	}

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "server.upload.path").String()
	basePath, _ := os.Getwd()
	if uploadPath == "" {
		// 不配置上传目录时，自动生成在运行目录下的public/upload
		uploadPath = path.Join(basePath, "./public/upload")
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			g.Log().Fatal(ctx, "File upload path configuration cannot be empty")
		}
	} else {
		uploadPath = path.Join(basePath, uploadPath)
	}
	s.AddStaticPath("/upload", uploadPath)

	// 开发阶段禁止浏览器缓存
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
		// 开发阶段才生成路由表
		rules(s)
	}
	// 设置默认语言
	g.I18n().SetLanguage("en")

	// 统一加载路由 和 路由中间件
	s.Group("/api", func(group *ghttp.RouterGroup) {
		// 加载
		router.LoadRouter(ctx, group)
	})

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("a critical error occurred:", err)
			}
		}()
		go func() {
			fmt.Println("Start Http serve, version:", gbuild.Info())
			s.Run()
			// 关闭
			stopSignal <- syscall.SIGQUIT
		}()
	}()

	return s
}

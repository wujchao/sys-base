package sys_base

import "github.com/gogf/gf/v2/os/gcmd"

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
	"os/signal"
	"sys-base/internal/cmd"

	_ "sys-base/internal/logic"
	"syscall"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var signalChannel = make(chan os.Signal, 1)
			// 启动服务
			cmd.RunServer(ctx, signalChannel)

			signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
			fmt.Println("Received shutdown signal:", <-signalChannel)

			fmt.Println("Successfully shut down the server")
			return nil
		},
	}
)

func Run() {
	Http.Run(gctx.GetInitCtx())
}

func RunService(ctx context.Context, stopSignal chan os.Signal) *ghttp.Server {
	if stopSignal == nil {
		stopSignal = make(chan os.Signal, 1)
	}
	return cmd.RunServer(ctx, stopSignal)
}

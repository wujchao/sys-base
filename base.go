package sys_base

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wujchao/sys-base/internal/cmd"
	"os"
	"os/signal"

	_ "github.com/wujchao/sys-base/internal/logic"
	"syscall"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var signalChannel = make(chan os.Signal, 1)
			s := g.Server()
			// 启动服务
			cmd.RunServer(ctx, s, signalChannel)

			signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
			fmt.Println("Received shutdown signal:", <-signalChannel)

			fmt.Println("Successfully shut down the server")
			return nil
		},
	}
)

func Run(ctx context.Context, s *ghttp.Server, stopSignal chan os.Signal) error {
	cmd.RunServer(ctx, s, stopSignal)
	return nil
}

func RunService(ctx context.Context, stopSignal chan os.Signal) *ghttp.Server {
	if stopSignal == nil {
		stopSignal = make(chan os.Signal, 1)
	}
	s := g.Server()
	cmd.RunServer(ctx, s, stopSignal)
	return s
}

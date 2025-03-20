package cmd

import (
	"context"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"
)

func RunSysPProf(stopSignal chan os.Signal) {
	pprofPort := g.Cfg().MustGet(context.Background(), "pprofPort").String()
	if pprofPort == "" {
		pprofPort = "8001"
	}
	runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // (非必需)开启对阻塞操作的跟踪
	err := statsviz.RegisterDefault()
	if err == nil {
		g.Log().Infof(context.Background(), "Point your browser to http://localhost:%s/debug/statsviz/", pprofPort)
		go func() {
			s := &http.Server{
				Addr:         ":" + pprofPort,
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 30 * time.Second,
				IdleTimeout:  30 * time.Second,
			}
			if err := s.ListenAndServe(); err != nil {
				fmt.Println("pprof server start error: ", err)
			}
			stopSignal <- syscall.SIGQUIT
		}()
	}
}

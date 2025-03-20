package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/useragent"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/internal/model/do"
	"sys-base/service"
	"sys-base/utility/gxid"
)

type sSysLoginLog struct {
}

func init() {
	service.RegisterSysLoginLog(&sSysLoginLog{})
}

func (s *sSysLoginLog) Add(ctx context.Context, params *model.LoginLogParams) {
	ua := useragent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.SysLoginLog{
		Id:        gxid.Gen(),
		LoginName: params.Username,
		Ipaddr:    params.Ip,
		Browser:   browser,
		Os:        ua.OS(),
		Status:    params.Status,
		Msg:       params.Msg,
		LoginTime: gtime.Now(),
		Module:    params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

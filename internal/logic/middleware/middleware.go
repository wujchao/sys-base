package middleware

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"regexp"
	"strings"
	"sys-base/consts"
	"sys-base/service"
	"sys-base/utility/gcasbin"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}

// Ctx 自定义上下文对象 服务端会忽略客户端的取消请求并继续往下执行
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	service.Context().Init(r)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// I18n 多语言支持
func (s *sMiddleware) I18n(r *ghttp.Request) {
	ctx := r.GetCtx()
	//g.Log().Debug(ctx, "测试接收到的数据：", r.GetBodyString())
	lang := r.GetHeader("Accept-Language")
	if lang == "" {
		lang = r.GetHeader("lang")
	}
	if lang == "zh-Hans" || lang == "zh" {
		lang = "zh-CN"
	}
	r.SetCtx(gi18n.WithLanguage(ctx, lang))
	r.Middleware.Next()
}

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	r.Middleware.Next()
}

// Permission 权限
func (s *sMiddleware) Permission(r *ghttp.Request) {
	re := regexp.MustCompile(consts.ApiVersionRegex)
	url := re.ReplaceAllString(r.RequestURI, "")
	e := gcasbin.GetCasbin()
	ok, err := e.Enforcer.Enforce("admin", url, r.Method)
	g.Dump(ok)
	if err != nil {
		r.Response.WriteJsonExit(errorRes(500, err.Error()))
		return
	}
	if !ok {
		r.Response.WriteJsonExit(errorRes(403, g.I18n().T(r.Context(), "Permission denied")))
		return
	}
	r.Middleware.Next()
}

// MiddlewareCORS 跨域处理
func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	//自定义跨域限制
	corsOptions := r.Response.DefaultCORSOptions()
	corsConfig := g.Cfg().MustGet(context.Background(), "server.allowedDomains").Strings()
	if corsConfig == nil || len(corsConfig) == 0 {
		//采用默认接受所有跨域
		r.Response.CORSDefault()
	} else {
		corsOptions.AllowDomain = corsConfig
		r.Response.CORS(corsOptions)
	}
	r.Middleware.Next()
}

// OperationLog 操作日志
func (s *sMiddleware) OperationLog(r *ghttp.Request) {
	r.Middleware.Next()
}

func (s *sMiddleware) MqttWebHook(r *ghttp.Request) {
	token := g.Cfg().MustGet(r.GetCtx(), "mqtt.token").String()
	reqToken := r.GetHeader("Authorization")
	if strings.HasPrefix(reqToken, "Bearer ") {
		reqToken = reqToken[7:]
	}
	if token != "" && token != reqToken {
		r.Response.WriteJsonExit(errorRes(403, g.I18n().T(r.Context(), "无权限")))
		r.ExitAll()
		return
	}
	g.Log().Info(r.GetCtx(), "MQTT WebHook body：", r.GetBodyString())
	r.Middleware.Next()
}

// Tracing 链路跟踪
func (s *sMiddleware) Tracing(r *ghttp.Request) {
	_, span := gtrace.NewSpan(r.Context(), r.Method+"_"+r.Request.RequestURI)
	defer span.End()
	r.Middleware.Next()
}

// DirectResponse CustomResponse 自定义直接返回
func (s *sMiddleware) DirectResponse(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	// 如果是直接返回
	_, ok := service.Context().GetData(r.GetCtx(), consts.DirectKey)
	g.Log().Debug(r.GetCtx(), "DirectResponse", ok)
	if ok {
		r.Response.WriteJsonExit(r.GetHandlerResponse())
		return
	}
}

// SecretKey校验方式
func secretKeyAuth(ctx context.Context, secretKey string) error {
	return errors.New("SecretKey is not supported")
}

func errorRes(code int, msg string) ghttp.DefaultHandlerResponse {
	if strings.HasPrefix(msg, "rpc error:") {
		descIndex := strings.Index(msg, "desc =")
		if descIndex > 0 {
			msg = msg[descIndex+7:]
		}
	}
	return ghttp.DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}

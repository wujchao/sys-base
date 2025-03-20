// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Ctx 自定义上下文对象 服务端会忽略客户端的取消请求并继续往下执行
		Ctx(r *ghttp.Request)
		// I18n 多语言支持
		I18n(r *ghttp.Request)
		// Auth 前台系统权限控制，用户必须登录才能访问
		Auth(r *ghttp.Request)
		// Permission 权限
		Permission(r *ghttp.Request)
		// MiddlewareCORS 跨域处理
		MiddlewareCORS(r *ghttp.Request)
		// OperationLog 操作日志
		OperationLog(r *ghttp.Request)
		MqttWebHook(r *ghttp.Request)
		// Tracing 链路跟踪
		Tracing(r *ghttp.Request)
		// DirectResponse CustomResponse 自定义直接返回
		DirectResponse(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}

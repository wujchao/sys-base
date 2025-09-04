// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/wujchao/sys-base/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		Init(r *ghttp.Request)
		// Get 获得上下文变量，如果没有设置，那么返回nil
		Get(ctx context.Context) *model.Context
		// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUser(ctx context.Context, ctxUser *model.ContextUser)
		// GetLoginUser 获取当前登陆用户信息
		GetLoginUser(ctx context.Context) *model.ContextUser
		// GetUserId 获取当前登录用户id
		GetUserId(ctx context.Context) *string
		// GetUserOrgId 获取当前登录用户企业ID
		GetUserOrgId(ctx context.Context) *string
		// GetUserName 获取当前登录用户账户
		GetUserName(ctx context.Context) string
		// UserIsManager 判断当前用户是否管理端
		UserIsManager(ctx context.Context) bool
		GetData(ctx context.Context, key string) (res interface{}, ok bool)
		SetData(ctx context.Context, key string, value interface{}) error
		SetAuthWay(ctx context.Context, requestWay string)
		GetAuthWay(ctx context.Context) string
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}

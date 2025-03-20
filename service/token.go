// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IToken interface {
		AddWhiteList(ctx context.Context, url string)
		GenerateToken(ctx context.Context, user any) (token string, err error)
		// RemoveToken 删除
		RemoveToken(ctx context.Context, token string) error
		// Exists Token是否存在
		Exists(ctx context.Context, token string) bool
		// RenewalToken 续期
		RenewalToken(ctx context.Context, token string) error
		// GetToken 获取Token 值
		GetToken(ctx context.Context) (token string)
		// GetCurUser 获取当前登录用户
		GetCurUser(ctx context.Context, user any) error
		GetCurUserMap(ctx context.Context) (map[string]interface{}, error)
		// GetUser 通过Token获取User, user为指针类型
		GetUser(ctx context.Context, token string, user any) error
		// TokenAuth Auth token中间件
		TokenAuth(r *ghttp.Request) error
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}

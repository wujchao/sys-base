package token

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"strings"
	"sys-base/service"
)

type sToken struct {
}

const tokenKey = "XX_TOKEN_KEY:"
const headerTokenKey = "Authorization"
const headerRefresh = "Refresh"

var WhiteList = map[string]struct{}{
	"/api/v1/login":    {},
	"/api/v1/logout":   {},
	"/api/v1/captcha":  {},
	"/api/v1/register": {},
	"/api/v1/refresh":  {},
}

func init() {
	service.RegisterToken(&sToken{})
}

func (s *sToken) AddWhiteList(ctx context.Context, url string) {
	WhiteList[url] = struct{}{}
}

func (s *sToken) GenerateToken(ctx context.Context, user any) (token string, err error) {
	token = guid.S()
	exp := g.Cfg().MustGet(ctx, "token.exp").Int64()
	if exp == 0 {
		exp = 86400
	}
	var userJson []byte

	switch v := user.(type) {
	case []byte:
		userJson = v
	case string:
		userJson = []byte(v)
	default:
		userJson, err = json.Marshal(v)
		if err != nil {
			return
		}
	}

	key := tokenKey + token
	_, err = g.Redis().Set(ctx, key, userJson, gredis.SetOption{
		TTLOption: gredis.TTLOption{EX: &exp},
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

// RemoveToken 删除
func (s *sToken) RemoveToken(ctx context.Context, token string) error {
	key := tokenKey + token
	_, err := g.Redis().Del(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

// Exists Token是否存在
func (s *sToken) Exists(ctx context.Context, token string) bool {
	key := tokenKey + token
	count, err := g.Redis().Exists(ctx, key)
	if err != nil || count == 0 {
		return false
	}
	// 检查成功，自动续刷新过期时间
	if g.Cfg().MustGet(ctx, "token.refresh", false).Bool() {
		_ = s.RenewalToken(ctx, token)
	}

	return true
}

// RenewalToken 续期
func (s *sToken) RenewalToken(ctx context.Context, token string) error {
	exp := g.Cfg().MustGet(ctx, "token.exp", 86400).Int64()
	_, err := g.Redis().Expire(ctx, token, exp)
	return err
}

// GetToken 获取Token 值
func (s *sToken) GetToken(ctx context.Context) (token string) {
	return strings.Trim(g.RequestFromCtx(ctx).Header.Get(headerTokenKey), " ")
}

// GetCurUser 获取当前登录用户
func (s *sToken) GetCurUser(ctx context.Context, user any) error {
	token := s.GetToken(ctx)
	if token == "" {
		return gerror.NewCode(gcode.CodeNotFound)
	}
	return s.GetUser(ctx, token, user)
}

func (s *sToken) GetCurUserMap(ctx context.Context) (map[string]interface{}, error) {
	user := map[string]interface{}{}
	err := s.GetCurUser(ctx, &user)
	return user, err
}

// GetUser 通过Token获取User, user为指针类型
func (s *sToken) GetUser(ctx context.Context, token string, user any) error {
	key := tokenKey + token
	v, err := g.Redis().Get(ctx, key)
	if err != nil {
		return err
	}
	if v.IsEmpty() {
		return gerror.NewCode(gcode.CodeNotAuthorized)
	}
	err = json.Unmarshal(v.Bytes(), user)
	return err
}

// TokenAuth Auth token中间件
func (s *sToken) TokenAuth(r *ghttp.Request) error {
	url := r.Request.URL.Path
	// 白名单
	if _, ok := WhiteList[url]; ok {
		return nil
	}

	token := strings.Trim(r.GetHeader(headerTokenKey), " ")
	// 空Token
	if token == "" {
		return gerror.NewCode(gcode.CodeNotFound)
	}
	// Toke不存在
	if has := s.Exists(r.Context(), token); !has {
		return gerror.NewCode(gcode.CodeNotAuthorized)
	}

	// 续期
	if strings.Trim(r.GetHeader(headerRefresh), " ") != "" {
		_ = s.RenewalToken(r.Context(), token)
	}

	return nil
}

package content

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sys-base/consts"
	"sys-base/internal/model"
	"sys-base/service"
)

type sContext struct{}

const ContextKey = "ContextKey"

func init() {
	service.RegisterContext(&sContext{})
}

func (s *sContext) Init(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 获取登录用户
	user := &model.ContextUser{}
	err := service.Token().GetCurUser(ctx, user)
	if err != nil {
		user = nil
	}

	contextModel := &model.Context{
		Data:    make(g.Map),
		AuthWay: "token",
		User:    user,
	}
	r.SetCtxVar(ContextKey, contextModel)

	return
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	ctxData := ctx.Value(ContextKey)
	if ctxData == nil {
		// 重新初始化
		r := g.RequestFromCtx(ctx)
		s.Init(r)
		return s.Get(r.GetCtx())
	}
	if localCtx, ok := ctxData.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	sysContext := s.Get(ctx)
	if sysContext == nil || sysContext.User == nil || sysContext.User.Id == "" {
		return nil
	}
	return sysContext.User
}

// GetUserId 获取当前登录用户id
func (s *sContext) GetUserId(ctx context.Context) *string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return &user.Id
	}
	return nil
}

// GetUserOrgId 获取当前登录用户企业ID
func (s *sContext) GetUserOrgId(ctx context.Context) *string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		// 数据错误
		if user.OrgId == "" && user.UserTypes == consts.UserTypeOrg {
			return nil
		}
		return &user.OrgId
	}
	return nil
}

// GetUserName 获取当前登录用户账户
func (s *sContext) GetUserName(ctx context.Context) string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.UserName
	}
	return ""
}

// UserIsManager 判断当前用户是否管理端
func (s *sContext) UserIsManager(ctx context.Context) bool {
	user := s.GetLoginUser(ctx)
	if user == nil {
		return false
	}
	if user.UserTypes != consts.UserTypeOrg && user.OrgId == "" {
		return true
	}
	return false
}

func (s *sContext) GetData(ctx context.Context, key string) (res interface{}, ok bool) {
	ctxData := s.Get(ctx)
	if ctxData == nil || ctxData.Data == nil || len(ctxData.Data) == 0 {
		return nil, false
	}
	res, ok = ctxData.Data[key]
	return
}

func (s *sContext) SetData(ctx context.Context, key string, value interface{}) error {
	ctxData := s.Get(ctx)
	if ctxData == nil {
		return gerror.NewCode(gcode.CodeInternalError)
	}
	if ctxData.Data == nil {
		ctxData.Data = make(g.Map)
	}
	ctxData.Data[key] = value
	s.Get(ctx).Data = ctxData.Data
	return nil
}

func (s *sContext) SetAuthWay(ctx context.Context, requestWay string) {
	ctxData := s.Get(ctx)
	if ctxData == nil {
		return
	}
	ctxData.AuthWay = requestWay
}

func (s *sContext) GetAuthWay(ctx context.Context) string {
	ctxData := s.Get(ctx)
	if ctxData == nil {
		return ""
	}
	return ctxData.AuthWay
}

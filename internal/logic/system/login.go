package system

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/wujchao/sys-base/api/sys_login/v1"
	"github.com/wujchao/sys-base/consts"
	"github.com/wujchao/sys-base/internal/dao"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/internal/model/entity"
	"github.com/wujchao/sys-base/service"
	"github.com/wujchao/sys-base/utility/cache"
	"time"
)

type sLogin struct {
}

func init() {
	service.RegisterLogin(&sLogin{})
}

func (s *sLogin) Login(ctx context.Context, in *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	// 错误次数限制
	if err = s.CheckPwdErrorCount(ctx, in.UserName); err != nil {
		return nil, err
	}
	// 判断验证码
	if in.VerifyKey != "" && !service.Captcha().VerifyString(in.VerifyKey, in.Captcha) {
		err = gerror.NewCode(consts.ParamsInvalidErrCode, g.I18n().T(ctx, "Verification code error"))
		return
	}

	// 登录失败处理
	loginStatus, loginMsg := 1, ""
	defer func() {
		s.loginAfter(ctx, in.UserName, loginStatus, loginMsg)
	}()

	// 获取用户
	userInfo, err := service.SysUsers().GetUserByUserNamePassword(ctx, in.UserName, in.Password)
	if err != nil {
		// 密码错误，登录失败
		loginStatus, loginMsg = 2, err.Error()
		return
	}

	if err = s.CheckStatus(ctx, userInfo); err != nil {
		// 用户信息错误，登录失败
		loginStatus, loginMsg = 3, err.Error()
		return nil, err
	}
	if err = s.CheckOrg(ctx, userInfo.OrgId); err != nil {
		// 租户信息错误，登录失败
		loginStatus, loginMsg = 4, err.Error()
		return nil, err
	}

	token, err := service.Token().GenerateToken(ctx, userInfo)
	if err != nil {
		return
	}
	res = &v1.LoginDoRes{
		UserInfo:    &model.LoginUserRes{},
		Token:       token,
		IsChangePwd: 0,
	}

	err = gconv.Struct(userInfo, res.UserInfo)
	if err != nil {
		loginStatus, loginMsg = 5, err.Error()
		return nil, err
	}
	// 更新用户信息
	_ = service.SysUsers().UpdateLoginInfo(ctx, &model.UserUpdateLoginInput{
		Id:            userInfo.Id,
		LastLoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
		LastLoginTime: gtime.Now(),
	})
	return
}

// CheckStatus 状态检测
func (s *sLogin) CheckStatus(ctx context.Context, user *entity.SysUsers) (err error) {
	if user.Status == 3 {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此帐户尚未验证"))
	}
	if user.Status != 1 {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此帐号已被禁用"))
	}
	// 有效期验证
	nowTime := gtime.Now()
	if user.EffectiveDate != nil && user.EffectiveDate.After(nowTime) {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此帐号未生效"))
	}
	if user.ExpirationDate != nil && user.ExpirationDate.Before(nowTime) {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此帐号已过期"))
	}
	return err
}

// CheckOrg 企业检测
func (s *sLogin) CheckOrg(ctx context.Context, orgId string) (err error) {
	if orgId == "" {
		return nil
	}
	org := new(entity.SysOrganization)
	if err = dao.SysOrganization.Ctx(ctx).Where("id", orgId).Scan(&org); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if org == nil || org.Id == "" {
		return gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "未找到此组织信息"))
	}

	if org.Status != 1 {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此组织已停用"))
	}

	membershipSetting := g.Cfg().MustGet(ctx, "service.membership", false).Bool()
	// 开启会员制，检查企业过期时间(过期可以登录，但不能访问设备数据）
	if membershipSetting {
		if org.ExpirationDate != nil && org.ExpirationDate.Before(gtime.Now()) {
			return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此组织已过期"))
		}
		if org.EffectiveDate != nil && org.EffectiveDate.After(gtime.Now()) {
			return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "此组织未生效"))
		}
	}
	return nil
}

func (s *sLogin) loginAfter(ctx context.Context, username string, status int, msg string) {
	key := consts.LoginErrorCount + username
	// 登录成功
	if status == 1 {
		// 清空错误次数
		_, _ = cache.Instance().Remove(ctx, key)
	} else {
		// 设置错误次数
		errDur := g.Cfg().MustGet(ctx, "service.login.errDur", 300).Duration()
		errNum := 0
		tempData, err := cache.Instance().Get(ctx, key)
		if err == nil {
			errNum = tempData.Int()
		}
		_ = cache.Instance().Set(ctx, key, errNum+1, errDur*time.Second)
	}
	// 获取IP地址 和 userAgent
	req := g.RequestFromCtx(ctx)
	ip := req.GetClientIp()
	userAgent := req.Header.Get("User-Agent")
	module := req.Header.Get("module")

	// 保存登录日志
	service.SysLoginLog().Add(ctx, &model.LoginLogParams{
		Status:    status,
		Username:  username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       msg,
		Module:    module,
	})
}

func (s *sLogin) CheckPwdErrorCount(ctx context.Context, username string) (err error) {
	tempData, err := cache.Instance().Get(ctx, consts.LoginErrorCount+username)
	errNum := 0
	if err == nil {
		errNum = tempData.Int()
	}
	maxErrorNum := g.Cfg().MustGet(ctx, "service.login.maxErrNum", 5).Int()
	errDur := g.Cfg().MustGet(ctx, "service.login.errDur", 300).Duration()
	if errNum > maxErrorNum {
		err = gerror.Newf("Too many incorrect password attempts, please try again after %d minutes", errDur/60)
	}

	return
}

func (s *sLogin) Logout(ctx context.Context) (err error) {
	token := service.Token().GetToken(ctx)
	if token == "" {
		return gerror.NewCode(consts.AuthExpiredErrCode, g.I18n().T(ctx, "未登录"))
	}
	err = service.Token().RemoveToken(ctx, token)
	return
}

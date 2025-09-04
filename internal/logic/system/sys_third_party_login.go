package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/consts"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/internal/model/entity"
	"sys-base/service"
	"sys-base/utility/gxid"
)

type sThirdPartyLogin struct {
}

func init() {
	service.RegisterThirdPartyLogin(&sThirdPartyLogin{})
}

func (s *sThirdPartyLogin) WxLogin(ctx context.Context, in *model.WechatLoginResponse) (out *model.WechatLoginOutput, err error) {
	// 微信登录
	thirdParty := &model.ThirdPartyLoginInput{
		ThirdPartyName:   consts.ThirdPartyWechat,
		ThirdPartyUserId: in.OpenID,
		AccessToken:      in.SessionKey,
	}
	data, err := s.Login(ctx, thirdParty)
	if err != nil {
		return
	}
	out = &model.WechatLoginOutput{
		ThirdPartyLoginOutput: data,
	}
	return
}

func (s *sThirdPartyLogin) Login(ctx context.Context, in *model.ThirdPartyLoginInput) (out *model.ThirdPartyLoginOutput, err error) {
	// 用户是否存在
	var thirdPartyUser *entity.SysThirdPartyUser
	err = dao.SysThirdPartyUser.Ctx(ctx).Where(dao.SysThirdPartyUser.Columns().ThirdPartyUserId, in.ThirdPartyUserId).
		Where(dao.SysThirdPartyUser.Columns().ThirdPartyName, in.ThirdPartyName).Scan(&thirdPartyUser)
	if err != nil {
		return nil, err
	}
	if thirdPartyUser == nil || thirdPartyUser.UserId == "" {
		// 注册
		userType := consts.UserTypeUser
		userName := in.ThirdPartyName + "_" + gxid.Gen()
		regParams := &model.ThirdPartyRegisterInput{
			UserName:         userName,
			ThirdPartyUserId: in.ThirdPartyUserId,
			AccessToken:      in.AccessToken,
			UserTypes:        userType,
		}
		out, err = s.Register(ctx, regParams)
		if err != nil {
			return
		}
	} else {
		// 登录
		return s.doLogin(ctx, thirdPartyUser.UserId)
	}
	return
}

func (s *sThirdPartyLogin) doLogin(ctx context.Context, userId string) (out *model.ThirdPartyLoginOutput, err error) {
	var userInfo *entity.SysUsers
	if err = dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().Id, userId).
		Scan(&userInfo); err != nil {
		return
	}
	if userInfo == nil {
		return nil, gerror.NewCode(consts.SystemErrCode)
	}
	// Do login
	if err = service.Login().CheckStatus(ctx, userInfo); err != nil {
		return
	}
	if err = service.Login().CheckOrg(ctx, userInfo.OrgId); err != nil {
		return
	}
	token, err := service.Token().GenerateToken(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	out = &model.ThirdPartyLoginOutput{
		Token:       token,
		UserInfo:    &model.LoginUserRes{},
		IsChangePwd: 0,
	}
	if err := gconv.Struct(userInfo, out.UserInfo); err != nil {
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

func (s *sThirdPartyLogin) Register(ctx context.Context, in *model.ThirdPartyRegisterInput) (out *model.ThirdPartyLoginOutput, err error) {
	// 创建用户
	nowTime := gtime.Now()
	user := &entity.SysUsers{
		Id:           gxid.Gen(),
		UserName:     in.UserName,
		UserTypes:    in.UserTypes,
		UserPassword: "", // 空密码
		IsAdmin:      consts.UserIsAdminYes,
		Status:       consts.UserStatusNormal,
		CreatedAt:    nowTime,
		UpdatedAt:    nowTime,
	}
	err = dao.SysUsers.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := dao.SysUsers.Ctx(ctx).Insert(user); err != nil {
			return err
		}
		// 关联第三方表
		thirdPartyUser := &entity.SysThirdPartyUser{
			UserId:           user.Id,
			ThirdPartyName:   consts.ThirdPartyWechat,
			ThirdPartyUserId: in.ThirdPartyUserId,
			AccessToken:      in.AccessToken,
			CreatedAt:        nowTime,
			UpdatedAt:        nowTime,
		}
		if _, err := dao.SysThirdPartyUser.Ctx(ctx).Insert(thirdPartyUser); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// 登录
	token, err := service.Token().GenerateToken(ctx, user)
	if err != nil {
		return
	}
	out = &model.ThirdPartyLoginOutput{
		Token:       token,
		UserInfo:    &model.LoginUserRes{},
		IsChangePwd: 0,
	}
	if err = gconv.Struct(user, out.UserInfo); err != nil {
		return
	}
	return
}

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/wujchao/sys-base/api/sys_login/v1"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ILogin interface {
		Login(ctx context.Context, in *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
		// CheckStatus 状态检测
		CheckStatus(ctx context.Context, user *entity.SysUsers) (err error)
		// CheckOrg 企业检测
		CheckOrg(ctx context.Context, orgId string) (err error)
		CheckPwdErrorCount(ctx context.Context, username string) (err error)
		Logout(ctx context.Context) (err error)
	}
	IRegister interface {
		Register(ctx context.Context, in *model.RegisterInput) (out *model.LoginOutput, err error)
	}
	ISysDoc interface {
		Read(ctx context.Context, key string) (out *model.SysDocOutput, err error)
	}
	ISysLoginLog interface {
		Add(ctx context.Context, params *model.LoginLogParams)
	}
	ISysMenu interface {
		ListOneself(ctx context.Context) (out []*model.SysMenuRes, err error)
		EncodeMenu(menus []*model.SysMenuRes, permission []string) (out []*model.SysMenuRes)
		GetTree(ctx context.Context) (res []*model.SysMenuRes, err error)
	}
	ISysRole interface {
		List(ctx context.Context, input *model.RoleListReqParam) (output *model.SysRoleListOutput, err error)
		Add(ctx context.Context, input *model.SysRoleInput) (output *model.SysRoleOutput, err error)
		Edit(ctx context.Context, in *model.SysRoleEditInput) (err error)
		Read(ctx context.Context, id string) (out *model.SysRoleOutput, err error)
		Delete(ctx context.Context, id string) (err error)
		GetRoleById(ctx context.Context, id string) (role *entity.SysRole, err error)
		GetRoleIdsByUserId(ctx context.Context, userId string) (roleIds []string, err error)
	}
	ISysRules interface {
		Tree(ctx context.Context) (res []*model.SysRulesTreeOutput, err error)
		ListOneself(ctx context.Context) (out []*model.SysRulesOutput, err error)
		GetRuleIdsByRoleIds(ctx context.Context, roleIds []string) (ruleIds []string, err error)
		// AddWithTx 事务关联角色权限
		AddWithTx(ctx context.Context, tx gdb.TX, ruleIds []string, roleId string) error
	}
	IThirdPartyLogin interface {
		WxLogin(ctx context.Context, in *model.WechatLoginResponse) (out *model.WechatLoginOutput, err error)
		Login(ctx context.Context, in *model.ThirdPartyLoginInput) (out *model.ThirdPartyLoginOutput, err error)
		Register(ctx context.Context, in *model.ThirdPartyRegisterInput) (out *model.ThirdPartyLoginOutput, err error)
	}
	ISysUsers interface {
		ListOneself(ctx context.Context, in *model.UserListParams) (out *model.UserListOutput, err error)
		ListWithRole(ctx context.Context, users []*model.UserOutput) (err error)
		Read(ctx context.Context, id string) (out *model.UserOutput, err error)
		Add(ctx context.Context, input *model.UserInput) (id string, err error)
		Edit(ctx context.Context, in *model.UserEditInput) (err error)
		GetUserByUsername(ctx context.Context, username string) (user *entity.SysUsers, err error)
		GetUserByUserNamePassword(ctx context.Context, username string, password string) (user *entity.SysUsers, err error)
		UpdateLoginInfo(ctx context.Context, in *model.UserUpdateLoginInput) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, id string) (err error)
		ChanPassSelf(ctx context.Context, in *model.UserSelfChanPasswordInput) (err error)
	}
)

var (
	localLogin           ILogin
	localRegister        IRegister
	localSysDoc          ISysDoc
	localSysLoginLog     ISysLoginLog
	localSysMenu         ISysMenu
	localSysRole         ISysRole
	localSysRules        ISysRules
	localThirdPartyLogin IThirdPartyLogin
	localSysUsers        ISysUsers
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

func Register() IRegister {
	if localRegister == nil {
		panic("implement not found for interface IRegister, forgot register?")
	}
	return localRegister
}

func RegisterRegister(i IRegister) {
	localRegister = i
}

func SysDoc() ISysDoc {
	if localSysDoc == nil {
		panic("implement not found for interface ISysDoc, forgot register?")
	}
	return localSysDoc
}

func RegisterSysDoc(i ISysDoc) {
	localSysDoc = i
}

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}

func SysRules() ISysRules {
	if localSysRules == nil {
		panic("implement not found for interface ISysRules, forgot register?")
	}
	return localSysRules
}

func RegisterSysRules(i ISysRules) {
	localSysRules = i
}

func ThirdPartyLogin() IThirdPartyLogin {
	if localThirdPartyLogin == nil {
		panic("implement not found for interface IThirdPartyLogin, forgot register?")
	}
	return localThirdPartyLogin
}

func RegisterThirdPartyLogin(i IThirdPartyLogin) {
	localThirdPartyLogin = i
}

func SysUsers() ISysUsers {
	if localSysUsers == nil {
		panic("implement not found for interface ISysUsers, forgot register?")
	}
	return localSysUsers
}

func RegisterSysUsers(i ISysUsers) {
	localSysUsers = i
}

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "sys-base/api/sys_login/v1"
	"sys-base/internal/model"
	"sys-base/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ILogin interface {
		Login(ctx context.Context, in *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
		CheckPwdErrorCount(ctx context.Context, username string) (err error)
		Logout(ctx context.Context) (err error)
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
	localLogin       ILogin
	localSysLoginLog ISysLoginLog
	localSysMenu     ISysMenu
	localSysRole     ISysRole
	localSysRules    ISysRules
	localSysUsers    ISysUsers
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

func SysUsers() ISysUsers {
	if localSysUsers == nil {
		panic("implement not found for interface ISysUsers, forgot register?")
	}
	return localSysUsers
}

func RegisterSysUsers(i ISysUsers) {
	localSysUsers = i
}

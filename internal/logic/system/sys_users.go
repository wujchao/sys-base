package system

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/consts"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/internal/model/do"
	"sys-base/internal/model/entity"
	"sys-base/service"
	"sys-base/utility"
	"sys-base/utility/gxid"
)

type sSysUsers struct {
}

func init() {
	service.RegisterSysUsers(&sSysUsers{})
}

func (s *sSysUsers) ListOneself(ctx context.Context, in *model.UserListParams) (out *model.UserListOutput, err error) {
	user := service.Context().GetLoginUser(ctx)
	out = new(model.UserListOutput)
	m := dao.SysUsers.Ctx(ctx)
	m = m.Where(dao.SysUsers.Columns().OrgId, user.OrgId)
	err = m.Handler(model.Pagination(&model.PaginationConfig{
		PaginationInput: in.PaginationInput,
	})).Handler(s.search(in)).ScanAndCount(&out.List, &out.Total, true)

	if err != nil {
		return nil, err
	}
	out.CurrentPage = in.Page
	// 用户角色
	err = s.ListWithRole(ctx, out.List)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysUsers) search(in *model.UserListParams) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if in.UserName != "" {
			m = m.WhereLike(dao.SysUsers.Columns().UserName, "%"+in.UserName+"%")
		}
		if in.Mobile != "" {
			m = m.WhereLike(dao.SysUsers.Columns().Mobile, "%"+in.Mobile+"%")
		}
		if in.Status != 0 {
			m = m.Where(dao.SysUsers.Columns().Status, in.Status)
		}
		return m
	}
}

func (s *sSysUsers) ListWithRole(ctx context.Context, users []*model.UserOutput) (err error) {
	if users == nil || len(users) == 0 {
		return nil
	}
	var userIds []string
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}

	var roleList []*model.SysRoleOutputWithUser
	err = dao.SysUsersRole.Ctx(ctx).As("ur").
		InnerJoin(dao.SysRole.Table(), "r", "r.id = ur.role_id").
		Fields("r.*", "ur.user_id").
		WhereIn("ur.user_id", userIds).
		Scan(&roleList)

	if err != nil {
		return err
	}
	if roleList == nil || len(roleList) == 0 {
		return nil
	}

	for k, user := range users {
		for _, role := range roleList {
			if user.Id == role.UserId {
				users[k].Roles = append(users[k].Roles, role.SysRoleOutput)
			}
		}
	}

	return nil
}
func (s *sSysUsers) Read(ctx context.Context, id string) (out *model.UserOutput, err error) {
	user := new(entity.SysUsers)
	err = dao.SysUsers.Ctx(ctx).Handler(service.Permiss().DataPermission(ctx)).
		Where(dao.SysUsers.Columns().Id, id).
		Scan(&user)

	if err != nil {
		return nil, gerror.NewCode(consts.SystemErrCode)
	}
	if user == nil {
		return nil, gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "用户不存在"))
	}
	if err = gconv.Scan(user, &out); err != nil {
		return nil, err
	}
	err = dao.SysRole.Ctx(ctx).
		WhereIn(dao.SysRole.Columns().Id, dao.SysUsersRole.Ctx(ctx).
			Fields(dao.SysUsersRole.Columns().RoleId).
			Where(dao.SysUsersRole.Columns().UserId, id)).
		Scan(&out.Roles)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	// 密码状态
	out.PassType = 1
	if user.UserPassword == "" {
		out.PassType = 2
	}
	return
}

func (s *sSysUsers) Add(ctx context.Context, input *model.UserInput) (id string, err error) {
	loginUser := service.Context().GetLoginUser(ctx)
	// 权限检查
	if err := s.checkAddPermissions(loginUser, &model.CheckPermissionsInput{
		OrgId:     input.OrgId,
		UserTypes: input.UserTypes,
		IsAdmin:   input.IsAdmin,
	}); err != nil {
		return "", err
	}

	// 判断用户是否存在
	exists, _ := dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().UserName, input.UserName).Exist()
	if exists {
		return "", gerror.NewCode(consts.DataExistErrCode, g.I18n().T(ctx, "账户已存在"))
	}

	// 创建用户密码
	password := utility.EncryptPassword(input.UserPassword, "")
	// 不设置用户类型时，默认类型
	if input.UserTypes == nil {
		adminUserType := consts.UserTypeAdmin
		orgUserType := consts.UserTypeOrg
		// 管理员默认为 1，企业默认为2 ，系统运维默认为9
		input.UserTypes = &orgUserType
		if input.OrgId == "" {
			input.UserTypes = &adminUserType
		}
	}
	// isAdmin默认为2
	if input.IsAdmin == 0 {
		input.IsAdmin = 2
	}
	// 开启事务
	err = dao.SysUsers.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id = gxid.Gen()
		_, err := dao.SysUsers.Ctx(ctx).TX(tx).OmitEmpty().Data(do.SysUsers{
			Id:           id,
			OrgId:        input.OrgId,
			UserName:     input.UserName,
			UserTypes:    input.UserTypes,
			Mobile:       input.Mobile,
			UserNickname: input.UserNickname,
			UserPassword: password,
			UserEmail:    input.UserEmail,
			Sex:          input.Sex,
			Avatar:       input.Avatar,
			DeptId:       input.DeptId,
			Remark:       input.Remark,
			IsAdmin:      input.IsAdmin,
			Address:      input.Address,
			Status:       1,
			CreatedBy:    loginUser.Id,
		}).Insert()

		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.NewCode(consts.DataExistErrCode, g.I18n().T(ctx, "添加用户失败"))
		}

		// 添加角色
		if input.RoleIds != nil {
			if err = s.userBindRole(ctx, tx, id, input.OrgId, input.RoleIds); err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sSysUsers) Edit(ctx context.Context, in *model.UserEditInput) (err error) {
	loginUser := service.Context().GetLoginUser(ctx)
	// 判断用户是否存在
	var editUser *entity.SysUsers
	err = dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().Id, in.Id).
		Scan(&editUser)
	if err != nil {
		return gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "用户不存在"))
	}
	// 权限检查
	if err := s.checkAddPermissions(loginUser, &model.CheckPermissionsInput{
		OrgId:     in.OrgId,
		UserTypes: nil,
		IsAdmin:   in.IsAdmin,
	}); err != nil {
		return err
	}

	updateUser := do.SysUsers{
		Id:           in.Id,
		Mobile:       in.Mobile,
		UserNickname: in.UserNickname,
		UserEmail:    in.UserEmail,
		Sex:          in.Sex,
		Avatar:       in.Avatar,
		Remark:       in.Remark,
		Address:      in.Address,
		Status:       in.Status,
		IsAdmin:      in.IsAdmin,
	}

	// 密码修改
	if in.UserPassword != "" {
		updateUser.UserPassword = utility.EncryptPassword(in.UserPassword, "")
	}

	// 开启事务
	err = dao.SysUsers.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新用户信息
		_, err = dao.SysUsers.Ctx(ctx).TX(tx).Data(updateUser).OmitEmptyData().Where(dao.SysUsers.Columns().Id, in.Id).Update()
		if err != nil {
			return err
		}
		// 如果有角色，先删除角色，再添加角色
		if in.RoleIds != nil {
			_, err = dao.SysUsersRole.Ctx(ctx).TX(tx).Where(dao.SysUsersRole.Columns().UserId, in.Id).Delete()
			if err != nil {
				return err
			}
			if err = s.userBindRole(ctx, tx, in.Id, in.OrgId, in.RoleIds); err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (s *sSysUsers) userBindRole(ctx context.Context, tx gdb.TX, userId string, orgId string, roleIds []string) (err error) {
	// 判断角色是否存在，且属于该企业
	var roleList []*entity.SysRole
	if err = dao.SysRole.Ctx(ctx).TX(tx).
		WhereIn(dao.SysRole.Columns().Id, roleIds).
		Where(dao.SysRole.Columns().OrgId, orgId).
		Where(dao.SysRole.Columns().Status, 1).Scan(&roleList); err != nil {
		g.Log().Error(ctx, err)
		return gerror.NewCode(consts.ParamsInvalidErrCode, g.I18n().T(ctx, "参数错误"))
	}

	if len(roleList) <= 0 {
		return nil
	}

	userRole := g.List{}
	for _, role := range roleList {
		if role.OrgId != orgId {
			return gerror.NewCode(consts.ParamsInvalidErrCode, g.I18n().T(ctx, "参数错误"))
		}
		userRole = append(userRole, map[string]interface{}{
			"userId": userId,
			"roleId": role.Id,
		})
	}
	if len(userRole) == 0 {
		return nil
	}
	// 绑定角色
	_, err = dao.SysUsersRole.Ctx(ctx).TX(tx).Data(userRole).Save()
	if err != nil {
		return err
	}

	return nil
}

func (s *sSysUsers) checkAddPermissions(actionUser *model.ContextUser, checkUser *model.CheckPermissionsInput) error {
	// 系统端
	if utility.IsSystemUser(actionUser) {
		// 系统端管理员
		if actionUser.IsAdmin == consts.UserIsAdminYes {
			return nil
		}
		// 系统端非管理员 不能添加系统端管理员
		if checkUser.OrgId == "" && checkUser.IsAdmin == consts.UserIsAdminYes {
			return gerror.NewCode(consts.PermsErrCode)
		}
		return nil
	}
	// 企业端
	// 1、不能添加系统端用户
	if checkUser.OrgId != "" || (checkUser.UserTypes != nil && *checkUser.UserTypes != consts.UserTypeOrg) {
		return gerror.NewCode(consts.ParamsInvalidErrCode)
	}
	// 2、非管理员不能添加管理员
	if actionUser.IsAdmin != consts.UserIsAdminYes && checkUser.IsAdmin == consts.UserIsAdminYes {
		return gerror.NewCode(consts.PermsErrCode)
	}

	return nil
}

func (s *sSysUsers) GetUserByUsername(ctx context.Context, username string) (user *entity.SysUsers, err error) {
	err = dao.SysUsers.Ctx(ctx).Fields(user).Where(dao.SysUsers.Columns().UserName, username).Scan(&user)
	if user == nil {
		err = gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "用户名不存在"))
		return
	}
	if user.Status == 2 {
		err = gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "此帐号已被禁用"))
		return
	}
	if user.Status == 3 {
		err = gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "此帐户尚未验证"))
	}
	return
}

func (s *sSysUsers) GetUserByUserNamePassword(ctx context.Context, username, password string) (user *entity.SysUsers, err error) {
	user, err = s.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "user", user)
	// 验证密码
	if !utility.PasswordVerify([]byte(password), []byte(user.UserPassword)) {
		return nil, gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "密码错误"))
	}
	return
}

func (s *sSysUsers) UpdateLoginInfo(ctx context.Context, in *model.UserUpdateLoginInput) (err error) {

	return err
}

// Delete 删除用户
func (s *sSysUsers) Delete(ctx context.Context, id string) (err error) {
	loginUser := service.Context().GetLoginUser(ctx)
	// 自已不能删除自己
	if loginUser.Id == id {
		return gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "不能删除自己"))
	}
	// 判断用户是否存在
	var editUser *entity.SysUsers
	err = dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().Id, id).
		Scan(&editUser)
	if err != nil || editUser == nil {
		return gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "数据不存在"))
	}

	// 系统平台端
	if utility.IsSystemUser(loginUser) {
		// 非管理员不能删除系统平台端的管理员
		if editUser.UserTypes != consts.UserTypeOrg {
			if editUser.IsAdmin == consts.UserIsAdminYes && loginUser.IsAdmin != consts.UserIsAdminYes {
				return gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "您无权限操作"))
			}
		}
	} else {
		// 企业端不可跨组织删除用户
		if editUser.OrgId != loginUser.OrgId {
			return gerror.NewCode(consts.PermsErrCode, g.I18n().T(ctx, "您无权限操作"))
		}
		// 非管理员不能删除管理员
		if editUser.IsAdmin == consts.UserIsAdminYes && loginUser.IsAdmin != consts.UserIsAdminYes {
			return gerror.New(g.I18n().T(ctx, "您无权限操作"))
		}
	}
	_, err = dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().Id, id).Delete()

	return err
}

func (s *sSysUsers) ChanPassSelf(ctx context.Context, in *model.UserSelfChanPasswordInput) (err error) {
	userId := service.Context().GetUserId(ctx)
	if userId == nil {
		return gerror.NewCode(consts.NotLoggedInErrCode)
	}
	var user *entity.SysUsers
	if err = dao.SysUsers.Ctx(ctx).Where("id", userId).Scan(&user); err != nil {
		return
	}

	if user.UserPassword != "" && !utility.PasswordVerify([]byte(in.OldPassword), []byte(user.UserPassword)) {
		return gerror.NewCode(consts.ParamsInvalidErrCode, g.I18n().T(ctx, "旧密码错误"))
	}
	_, err = dao.SysUsers.Ctx(ctx).Where("id", userId).Data(g.Map{
		dao.SysUsers.Columns().UserPassword: utility.EncryptPassword(in.Password, ""),
	}).Update()
	return
}

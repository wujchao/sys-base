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
	"sys-base/utility/gxid"
)

type sSysRole struct {
}

func init() {
	service.RegisterSysRole(&sSysRole{})
}

func (s *sSysRole) List(ctx context.Context, input *model.RoleListReqParam) (output *model.SysRoleListOutput, err error) {
	loginUser := service.Context().GetLoginUser(ctx)
	if loginUser.UserTypes == consts.UserTypeOrg || loginUser.OrgId != "" {
		input.OrgId = loginUser.OrgId
	}

	output = new(model.SysRoleListOutput)

	m := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().OrgId, input.OrgId)
	err = m.Handler(model.Pagination(&model.PaginationConfig{
		PaginationInput: input.PaginationInput,
		Fields:          nil,
	})).Handler(s.search(input)).ScanAndCount(&output.List, &output.Total, true)
	if err != nil {
		return nil, err
	}
	output.CurrentPage = input.Page
	g.Log().Debug(ctx, "output", output)
	return
}

func (s *sSysRole) search(input *model.RoleListReqParam) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if input.Status != 0 {
			m = m.Where(dao.SysRole.Columns().Status, input.Status)
		}
		if input.Name != "" {
			m = m.WhereLike(dao.SysRole.Columns().Name, "%"+input.Name+"%")
		}
		return m
	}
}

func (s *sSysRole) Add(ctx context.Context, input *model.SysRoleInput) (output *model.SysRoleOutput, err error) {
	loginUser := service.Context().GetLoginUser(ctx)

	// 数据限定
	m := dao.SysRole.Ctx(ctx)
	if loginUser.UserTypes == consts.UserTypeOrg || loginUser.OrgId != "" {
		m = m.Where(dao.SysRole.Columns().OrgId, loginUser.OrgId)
	}
	// 判断角色是否已存在
	exists, _ := m.Where(dao.SysRole.Columns().Name, input.Name).Exist()
	if exists {
		return nil, gerror.NewCode(consts.DataExistErrCode)
	}
	output = new(model.SysRoleOutput)
	if err = gconv.Scan(input, &output); err != nil {
		return nil, err
	}
	output.RoleType = uint(loginUser.UserTypes)

	// 判断上级角色是否存在
	if input.ParentId != "" {
		parentRole, err := s.GetRoleById(ctx, input.ParentId)
		if parentRole == nil || err != nil {
			return nil, gerror.NewCode(consts.ParamsInvalidErrCode, g.I18n().T(ctx, "参数错误"))
		}
		output.Path = parentRole.Path + "," + gconv.String(parentRole.Id)
	}

	// 开启事务
	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		output.Id = gxid.Gen()
		_, err := dao.SysRole.Ctx(ctx).TX(tx).Data(do.SysRole{
			Id:        output.Id,
			OrgId:     loginUser.OrgId,
			ParentId:  output.ParentId,
			RoleType:  output.RoleType,
			Path:      output.Path,
			ListOrder: output.ListOrder,
			Name:      output.Name,
			Remark:    output.Remark,
			Status:    output.Status,
			CreatedBy: loginUser.Id,
		}).Insert()
		if err != nil {
			return err
		}
		// 关联权限
		if err = service.SysRules().AddWithTx(ctx, tx, input.Rules, output.Id); err != nil {
			return err
		}
		return nil
	})

	return
}

func (s *sSysRole) Edit(ctx context.Context, in *model.SysRoleEditInput) (err error) {
	loginOrgId := service.Context().GetUserOrgId(ctx)
	if loginOrgId == nil {
		return err
	}
	// 判断 是否有修改角色权限
	m := dao.SysRole.Ctx(ctx)
	if *loginOrgId != "" {
		m = m.Where(dao.SysRole.Columns().OrgId, loginOrgId)
	}
	m = m.Where(dao.SysRole.Columns().Id, in.Id)
	num, _ := m.Count()
	if num <= 0 {
		return gerror.NewCode(consts.DataNotFoundErrCode)
	}
	// 开启事务
	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 修改角色
		_, err = m.TX(tx).Data(in).Update()
		if err != nil {
			return err
		}
		// 如果有授权，修改授权
		if len(in.Rules) > 0 {
			_, err = dao.SysRoleRules.Ctx(ctx).TX(tx).Where(dao.SysRoleRules.Columns().RoleId, in.Id).Delete()
			if err != nil {
				return err
			}
			if err = service.SysRules().AddWithTx(ctx, tx, in.Rules, in.Id); err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func (s *sSysRole) Read(ctx context.Context, id string) (out *model.SysRoleOutput, err error) {
	loginOrgId := service.Context().GetUserOrgId(ctx)
	if loginOrgId == nil {
		return nil, err
	}
	m := dao.SysRole.Ctx(ctx)
	if *loginOrgId != "" {
		m = m.Where(dao.SysRole.Columns().OrgId, loginOrgId)
	}
	if err = m.Where(dao.SysRole.Columns().Id, id).Scan(&out); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return
}

func (s *sSysRole) Delete(ctx context.Context, id string) (err error) {
	loginOrgId := service.Context().GetUserOrgId(ctx)
	if loginOrgId == nil {
		return
	}
	m := dao.SysRole.Ctx(ctx)
	if "" != *loginOrgId {
		m = m.Where(dao.SysRole.Columns().OrgId, loginOrgId)
	}
	_, err = m.Where(dao.SysRole.Columns().Id, id).Delete()
	return
}

func (s *sSysRole) GetRoleById(ctx context.Context, id string) (role *entity.SysRole, err error) {

	if err = dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, id).
		Where(dao.SysRole.Columns().Status, 1).
		Scan(&role); err != nil {
		return nil, gerror.NewCode(consts.SystemErrCode)
	}
	return
}

func (s *sSysRole) GetRoleIdsByUserId(ctx context.Context, userId string) (roleIds []string, err error) {
	v, err := dao.SysUsersRole.Ctx(ctx).Fields(dao.SysUsersRole.Columns().RoleId).
		Where(dao.SysUsersRole.Columns().UserId, userId).Array()
	if err != nil {
		return nil, err
	}
	if v == nil || len(v) == 0 {
		return nil, nil
	}
	err = gconv.Scan(v, &roleIds)
	if err != nil {
		return nil, gerror.Wrap(err, "Failed to convert role IDs")
	}
	return roleIds, nil
}

package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sys-base/consts"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/internal/model/entity"
	"sys-base/service"
	"sys-base/utility"
	"sys-base/utility/gxid"
)

type sRegister struct {
}

func init() {
	service.RegisterRegister(&sRegister{})
}

func (s *sRegister) Register(ctx context.Context, in *model.RegisterInput) (out *model.LoginOutput, err error) {
	exist, err := dao.SysUsers.Ctx(ctx).Where(dao.SysUsers.Columns().UserName, in.UserName).Exist()
	if err != nil {
		return
	}
	if exist {
		err = gerror.NewCode(consts.DataExistErrCode, g.I18n().T(ctx, "账户已存在"))
		return
	}
	in.UserPassword = utility.EncryptPassword(in.UserPassword, "")
	// 写入用户
	var user entity.SysUsers
	if err = gconv.Scan(in, &user); err != nil {
		return
	}
	user.Id = gxid.Gen()
	user.CreatedAt = gtime.New()
	user.UpdatedAt = gtime.New()
	user.DeletedAt = 0
	if _, err = dao.SysUsers.Ctx(ctx).Insert(user); err != nil {
		return
	}
	// 登录
	token, err := service.Token().GenerateToken(ctx, user)
	if err != nil {
		return
	}
	out = &model.LoginOutput{
		Token:       token,
		UserInfo:    &model.LoginUserRes{},
		IsChangePwd: 0,
	}
	if err = gconv.Struct(user, out.UserInfo); err != nil {
		return
	}
	return
}

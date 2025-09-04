// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_user

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_user/v1"
)

type ISysUserV1 interface {
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserRead(ctx context.Context, req *v1.UserReadReq) (res *v1.UserReadRes, err error)
	UserAdd(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error)
	UserEdit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserEditRes, err error)
	UserDeleteById(ctx context.Context, req *v1.UserDeleteByIdReq) (res *v1.UserDeleteByIdRes, err error)
	ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error)
}

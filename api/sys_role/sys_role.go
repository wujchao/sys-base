// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_role

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_role/v1"
)

type ISysRoleV1 interface {
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	RoleRead(ctx context.Context, req *v1.RoleReadReq) (res *v1.RoleReadRes, err error)
	RoleAdd(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error)
	RoleEdit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleEditRes, err error)
	RoleDeleteById(ctx context.Context, req *v1.RoleDeleteByIdReq) (res *v1.RoleDeleteByIdRes, err error)
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_admin

import (
	"context"

	"sys-base/api/sys_admin/v1"
)

type ISysAdminV1 interface {
	AdminList(ctx context.Context, req *v1.AdminListReq) (res *v1.AdminListRes, err error)
	AdminRead(ctx context.Context, req *v1.AdminReadReq) (res *v1.AdminReadRes, err error)
	AdminAdd(ctx context.Context, req *v1.AdminAddReq) (res *v1.AdminAddRes, err error)
	AdminEdit(ctx context.Context, req *v1.AdminEditReq) (res *v1.AdminEditRes, err error)
	AdminDeleteById(ctx context.Context, req *v1.AdminDeleteByIdReq) (res *v1.AdminDeleteByIdRes, err error)
}

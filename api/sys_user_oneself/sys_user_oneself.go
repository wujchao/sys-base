// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_user_oneself

import (
	"context"

	"sys-base/api/sys_user_oneself/v1"
)

type ISysUserOneselfV1 interface {
	SelfUserRead(ctx context.Context, req *v1.SelfUserReadReq) (res *v1.SelfUserReadRes, err error)
	SelfUserEdit(ctx context.Context, req *v1.SelfUserEditReq) (res *v1.SelfUserEditRes, err error)
	SelfUserChanPass(ctx context.Context, req *v1.SelfUserChanPassReq) (res *v1.SelfUserChanPassRes, err error)
}

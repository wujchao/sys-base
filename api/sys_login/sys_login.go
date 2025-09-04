// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_login

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_login/v1"
)

type ISysLoginV1 interface {
	LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
	LoginOut(ctx context.Context, req *v1.LoginOutReq) (res *v1.LoginOutRes, err error)
}

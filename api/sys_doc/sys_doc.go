// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_doc

import (
	"context"

	"sys-base/api/sys_doc/v1"
)

type ISysDocV1 interface {
	SysDocRead(ctx context.Context, req *v1.SysDocReadReq) (res *v1.SysDocReadRes, err error)
}

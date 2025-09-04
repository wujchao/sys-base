// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_oneself

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_oneself/v1"
)

type ISysOneselfV1 interface {
	OneselfMenu(ctx context.Context, req *v1.OneselfMenuReq) (res *v1.OneselfMenuRes, err error)
	OneselfRulesTree(ctx context.Context, req *v1.OneselfRulesTreeReq) (res *v1.OneselfRulesTreeRes, err error)
}

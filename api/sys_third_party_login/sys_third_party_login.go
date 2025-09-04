// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_third_party_login

import (
	"context"

	"sys-base/api/sys_third_party_login/v1"
)

type ISysThirdPartyLoginV1 interface {
	WxLogin(ctx context.Context, req *v1.WxLoginReq) (res *v1.WxLoginRes, err error)
}

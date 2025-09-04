// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_setting

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_setting/v1"
)

type ISysSettingV1 interface {
	SettingAdd(ctx context.Context, req *v1.SettingAddReq) (res *v1.SettingAddRes, err error)
	SettingGetById(ctx context.Context, req *v1.SettingGetByIdReq) (res *v1.SettingGetByIdRes, err error)
	SettingGetByGroup(ctx context.Context, req *v1.SettingGetByGroupReq) (res *v1.SettingGetByGroupRes, err error)
	SettingDeleteById(ctx context.Context, req *v1.SettingDeleteByIdReq) (res *v1.SettingDeleteByIdRes, err error)
}

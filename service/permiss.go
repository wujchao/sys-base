// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IPermiss interface {
		DataPermission(ctx context.Context) func(m *gdb.Model) *gdb.Model
	}
)

var (
	localPermiss IPermiss
)

func Permiss() IPermiss {
	if localPermiss == nil {
		panic("implement not found for interface IPermiss, forgot register?")
	}
	return localPermiss
}

func RegisterPermiss(i IPermiss) {
	localPermiss = i
}

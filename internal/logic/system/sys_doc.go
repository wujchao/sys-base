package system

import (
	"context"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/service"
)

type sSysDoc struct {
}

func init() {
	service.RegisterSysDoc(&sSysDoc{})
}

func (s *sSysDoc) Read(ctx context.Context, key string) (out *model.SysDocOutput, err error) {
	err = dao.SysDoc.Ctx(ctx).
		Where(dao.SysDoc.Columns().Key, key).
		OrderDesc(dao.SysDoc.Columns().UpdatedAt).
		Scan(&out)
	return
}

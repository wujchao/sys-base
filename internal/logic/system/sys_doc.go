package system

import (
	"context"
	"github.com/wujchao/sys-base/internal/dao"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"
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

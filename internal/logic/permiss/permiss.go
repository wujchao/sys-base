package permiss

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"sys-base/service"
	"sys-base/utility"
)

type sPermiss struct {
}

func init() {
	service.RegisterPermiss(&sPermiss{})
}
func (s *sPermiss) DataPermission(ctx context.Context) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		user := service.Context().GetLoginUser(ctx)
		if user == nil || user.Id == "" {
			return m
		}
		if utility.IsSystemUser(user) {
			return m
		}
		return m.Where("org_id = ?", user.OrgId)
	}
}

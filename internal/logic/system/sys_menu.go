package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"sys-base/consts"
	"sys-base/internal/dao"
	"sys-base/internal/model"
	"sys-base/service"
	"sys-base/utility"
)

type sSysMenu struct {
}

func init() {
	service.RegisterSysMenu(&sSysMenu{})
}

func (s *sSysMenu) ListOneself(ctx context.Context) (out []*model.SysMenuRes, err error) {
	user := service.Context().GetLoginUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(consts.NotLoggedInErrCode)
	}
	// 查询条件
	m := dao.SysMenu.Ctx(ctx).WhereNot(dao.SysMenu.Columns().IsHide, 1)
	// 权限
	var permission []string
	// 管理端
	if user.UserTypes == 1 && user.OrgId == "" {
		m = m.WhereIn(dao.SysMenu.Columns().ModuleType, []int{1, 2})
		permission = append(permission, "super")
	} else if user.UserTypes == 9 && user.OrgId == "" {
		// 系统端
		m = m.WhereIn(dao.SysMenu.Columns().ModuleType, []int{1, 2, 0})
		permission = append(permission, "system")
	} else {
		m = m.WhereIn(dao.SysMenu.Columns().ModuleType, []int{1, 3})
		permission = append(permission, "org")
	}

	// 非超级管理员
	if user.IsAdmin != consts.UserIsAdminYes {
		// 获取授权列表
		rules, err := service.SysRules().ListOneself(ctx)
		if err != nil {
			return nil, err
		}
		if rules == nil || len(rules) == 0 {
			return make([]*model.SysMenuRes, 0), nil
		}
		// 接口
		var apiPaths []string
		// 菜单
		var menuIds []string
		for _, rule := range rules {
			if rule.Menus != "" {
				menuIds = append(menuIds, gconv.SliceStr(rule.Menus)...)
			}
			if rule.Apis != "" {
				apiPaths = append(apiPaths, gconv.SliceStr(rule.Apis)...)
			}
		}

		err = utility.SetPermissionForUid(user.Id, apiPaths)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.NewCode(consts.SystemErrCode)
		}

		if len(menuIds) == 0 {
			return make([]*model.SysMenuRes, 0), nil
		}

		m = m.WhereIn(dao.SysMenu.Columns().Id, menuIds).
			WhereOr(dao.SysMenu.Columns().Open, 1)
	}

	err = m.Scan(&out)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	out = s.EncodeMenu(out, permission)
	return
}

func (s *sSysMenu) EncodeMenu(menus []*model.SysMenuRes, permission []string) (out []*model.SysMenuRes) {
	if menus == nil || len(menus) == 0 {
		return menus
	}
	for k, v := range menus {
		menus[k].Meta = &model.SysMenuMeta{
			Title:      v.Title,
			Icon:       v.Icon,
			Hidden:     utility.StatusToBool(v.Hidden),
			Affix:      utility.StatusToBool(v.Affix),
			NoTagsView: utility.StatusToBool(v.NoTagsView),
			IsCached:   utility.StatusToBool(v.IsCached),
			Permission: permission,
		}
	}
	return menus
}

func (s *sSysMenu) GetTree(ctx context.Context) (res []*model.SysMenuRes, err error) {
	menuList, err := s.ListOneself(ctx)
	if err != nil {
		return nil, err
	}
	if menuList == nil {
		return nil, nil
	}
	// 获取所有根节点
	var parentNodes []*model.SysMenuRes
	for _, v := range menuList {
		if v.ParentId == "" {
			var parentNode *model.SysMenuRes
			if err = gconv.Scan(v, &parentNode); err != nil {
				return
			}
			parentNodes = append(parentNodes, parentNode)
		}
	}
	// 排序
	sort.SliceStable(parentNodes, func(i, j int) bool {
		return parentNodes[i].Weigh > parentNodes[j].Weigh
	})
	// buildTree
	res = buildTree(parentNodes, menuList)
	return
}

func buildTree(parentNode []*model.SysMenuRes, data []*model.SysMenuRes) (dataTree []*model.SysMenuRes) {
	// 遍历所有根节点
	for k, parent := range parentNode {
		// 查询根节点下的所有子节点
		for _, children := range data {
			var node *model.SysMenuRes
			if parent.Id == children.ParentId {
				if err := gconv.Scan(children, &node); err != nil {
					return
				}
				parentNode[k].Children = append(parentNode[k].Children, node)
			}
		}
		// 对子节点进行排序
		sort.SliceStable(parent.Children, func(i, j int) bool {
			return parent.Children[i].Weigh > parent.Children[j].Weigh
		})
		// 子节点递归
		buildTree(parent.Children, data)
	}
	return parentNode
}

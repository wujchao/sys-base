package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wujchao/sys-base/consts"
	"github.com/wujchao/sys-base/internal/dao"
	"github.com/wujchao/sys-base/internal/model"
	"github.com/wujchao/sys-base/service"
	"sort"
)

type sSysRules struct {
}

func (s *sSysRules) Tree(ctx context.Context) (res []*model.SysRulesTreeOutput, err error) {
	list, err := s.ListOneself(ctx)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, nil
	}

	var treeList []*model.SysRulesTreeOutput
	if err = gconv.Scan(list, &treeList); err != nil {
		return nil, err
	}
	// 国际化
	for k, v := range treeList {
		treeList[k].Name = g.I18n().T(ctx, v.Name)
	}

	// 获取所有根节点
	var parentNodes []*model.SysRulesTreeOutput
	for _, v := range treeList {
		if v.ParentId == "" {
			var parentNode *model.SysRulesTreeOutput
			if err = gconv.Scan(v, &parentNode); err != nil {
				return nil, err
			}
			parentNodes = append(parentNodes, parentNode)
		}
	}

	// 排序
	sort.SliceStable(parentNodes, func(i, j int) bool {
		return parentNodes[i].ListOrder < parentNodes[j].ListOrder
	})

	res = s.buildTree(parentNodes, treeList)

	return
}

func (s *sSysRules) buildTree(parentNode []*model.SysRulesTreeOutput, data []*model.SysRulesTreeOutput) (dataTree []*model.SysRulesTreeOutput) {
	// 遍历所有根节点
	for k, parent := range parentNode {
		// 查询根节点下所有子节点
		for _, children := range data {
			var node *model.SysRulesTreeOutput
			if parent.Id == children.ParentId {
				if err := gconv.Scan(children, &node); err != nil {
					return nil
				}
				parentNode[k].Children = append(parentNode[k].Children, node)
			}
		}
		// 对子节点进行排序
		sort.SliceStable(parent.Children, func(i, j int) bool {
			return parent.Children[i].ListOrder < parent.Children[j].ListOrder
		})

		// 子节点递归
		s.buildTree(parent.Children, data)
	}
	return parentNode
}

func (s *sSysRules) ListOneself(ctx context.Context) (out []*model.SysRulesOutput, err error) {
	user := service.Context().GetLoginUser(ctx)
	m := dao.SysRules.Ctx(ctx)

	if user == nil || user.Id == "" {
		return nil, gerror.NewCode(consts.ParamsInvalidErrCode)
	}

	// 超级管理员
	if user.IsAdmin == 1 {
		var moduleTypes []uint64
		if user.OrgId == "" && user.UserTypes != 2 {
			moduleTypes = []uint64{1, 2}
		} else {
			moduleTypes = []uint64{1, 3}
		}
		err = m.WhereIn(dao.SysRules.Columns().ModuleType, moduleTypes).
			Where(dao.SysRules.Columns().Status, 1).
			OrderDesc(dao.SysRules.Columns().ListOrder).
			Scan(&out)
		return
	}

	// 非超级管理
	roleIds, err := service.SysRole().GetRoleIdsByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	if roleIds == nil || len(roleIds) == 0 {
		return nil, gerror.NewCode(consts.DataNotFoundErrCode, g.I18n().T(ctx, "User role does not exist"))
	}

	// 通过角色获取权限
	ruleIds, err := s.GetRuleIdsByRoleIds(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	if ruleIds == nil || len(ruleIds) == 0 {
		return nil, nil
	}
	err = dao.SysRules.Ctx(ctx).WhereIn(dao.SysRules.Columns().Id, ruleIds).Scan(&out)
	return
}

func (s *sSysRules) GetRuleIdsByRoleIds(ctx context.Context, roleIds []string) (ruleIds []string, err error) {
	res, err := dao.SysRoleRules.Ctx(ctx).Fields(dao.SysRoleRules.Columns().RulesId).
		WhereIn(dao.SysRoleRules.Columns().RoleId, roleIds).Array()
	if err != nil {
		return nil, err
	}
	if res == nil || len(res) == 0 {
		return nil, nil
	}
	err = gconv.Scan(res, &ruleIds)
	if err != nil {
		return nil, gerror.Wrap(err, "Failed to convert rule IDs")
	}
	return
}

// AddWithTx 事务关联角色权限
func (s *sSysRules) AddWithTx(ctx context.Context, tx gdb.TX, ruleIds []string, roleId string) error {
	if len(ruleIds) == 0 || roleId == "" {
		return nil
	}
	if ok := s.exist(ctx, ruleIds); !ok {
		return gerror.NewCode(consts.ParamsInvalidErrCode)
	}
	roleRules := g.List{}
	for _, rid := range ruleIds {
		roleRules = append(roleRules, map[string]interface{}{
			dao.SysRoleRules.Columns().RoleId:  roleId,
			dao.SysRoleRules.Columns().RulesId: rid,
		})
	}
	_, err := dao.SysRoleRules.Ctx(ctx).TX(tx).Data(roleRules).Save()
	if err != nil {
		return err
	}
	return nil
}

func (s *sSysRules) exist(ctx context.Context, ruleIds []string) bool {
	res, err := dao.SysRules.Ctx(ctx).Fields("id").
		WhereIn(dao.SysRules.Columns().Id, ruleIds).
		Where(dao.SysRules.Columns().Status, 1).
		Array()
	if err != nil {
		return false
	}
	resMap := map[string]bool{}
	for _, v := range res {
		resMap[v.String()] = true
	}

	for _, v := range ruleIds {
		if ok := resMap[v]; !ok {
			return false
		}
	}
	return true
}

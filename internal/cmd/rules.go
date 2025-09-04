package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/wujchao/sys-base/internal/dao"
	"github.com/wujchao/sys-base/utility"
	"reflect"
	"strings"
)

type rule struct {
	Id         string   `json:"id"         orm:"id"          description:"Permission identifier, globally unique"`                 // Permission identifier, globally unique
	ParentId   string   `json:"parentId"   orm:"parent_id"   description:"ParentId"`                                               // ParentId
	Name       string   `json:"name"       orm:"name"        description:"Name"`                                                   // Name
	Apis       []string `json:"apis"       orm:"apis"        description:"List of API addresses"`                                  // List of API addresses
	Menus      []string `json:"menus"      orm:"menus"       description:"Menus displayed corresponding to the permission"`        // Menus displayed corresponding to the permission
	ListOrder  uint     `json:"listOrder"  orm:"list_order"  description:"Sorting, from smallest to largest"`                      // Sorting, from smallest to largest
	FrontIds   []string `json:"frontIds"   orm:"front_ids"   description:"pre-permission"`                                         // pre-permission
	ModuleType int      `json:"moduleType" orm:"module_type" description:"Type: 1 for all, 2 for management, 3 for enterprise"`    // Type: 1 for all, 2 for management, 3 for enterprise
	AuthType   uint     `json:"authType"   orm:"auth_type"   description:"Authorization Type, 1 Authorization, 2 Login, 3 Public"` // Authorization Type, 1 Authorization, 2 Login, 3 Public
	Status     uint     `json:"status"     orm:"status"      description:"Status"`                                                 // Status
	Desc       string   `json:"desc"       orm:"desc"        description:"Description"`                                            // Description
	Remark     string   `json:"remark"     orm:"remark"      description:"Remarks"`                                                // Remarks
}

func rules(s *ghttp.Server) {
	path := gfile.Join(gfile.Pwd(), "routes.json")
	if gfile.Exists(path) {
		return
	}

	go func() {
		// 等待启动完成
		for {
			if s.Status() == 1 {
				break
			}
		}
		ruleMap := map[string]rule{}

		for _, item := range s.GetRoutes() {
			switch item.Type {
			case "middleware", "hook":
				continue
			}
			if item.Handler.Info.IsStrictRoute {
				if item.Handler.Info.Type == nil || item.Handler.Info.Type.NumIn() != 2 {
					continue
				}
				ObjReq := reflect.New(item.Handler.Info.Type.In(1))
				metaData := gmeta.Data(ObjReq)
				ruleStr, ok := metaData["rules"]
				if !ok || ruleStr == "" {
					continue
				}

				ids := strings.Fields(ruleStr)
				authType := uint(1)
				if metaData["auth"] != "" {
					switch metaData["auth"] {
					case "login", "2":
						authType = 2
						break
					case "open", "3":
						authType = 3
					}
				}
				module := 1
				if metaData["module"] != "" {
					switch metaData["module"] {
					case "all", "1":
						break
					case "admin", "2":
						module = 2
						break
					case "org", "3":
						module = 3
						break
					}
				}
				for _, id := range ids {
					if id == "" {
						continue
					}
					if existingRule, ok := ruleMap[id]; ok {
						existingRule.Apis = append(existingRule.Apis, utility.EncodeRuleApi(item))
						if metaData["menu"] != "" {
							if existingRule.Menus == nil {
								existingRule.Menus = []string{}
							}
							existingRule.Menus = append(existingRule.Menus, strings.Fields(metaData["menu"])...)
						}
						ruleMap[id] = existingRule
					} else {
						ruleItem := rule{
							Id:         id,
							ParentId:   metaData["parent"],
							Name:       metaData["name"],
							Apis:       []string{utility.EncodeRuleApi(item)},
							Menus:      strings.Fields(metaData["menu"]),
							FrontIds:   strings.Fields(metaData["front"]),
							ListOrder:  100,
							Status:     1,
							AuthType:   authType,
							ModuleType: module,
							Desc:       metaData["desc"],
							Remark:     metaData["summary"],
						}
						ruleMap[id] = ruleItem
					}
				}
			}
		}
		_, err := g.DB().Exec(context.TODO(), "TRUNCATE TABLE "+dao.SysRules.Table())
		if err != nil {
			g.Log().Error(context.TODO(), err)
			return
		}
		err = insertRules(ruleMap)
		if err != nil {
			g.Log().Error(context.TODO(), err)
			return
		}
		err = gfile.PutContents(path, gconv.String(ruleMap))
		if err != nil {
			g.Log().Error(context.TODO(), err)
			return
		}
	}()
}

// []string{} 去重
func unique(slice []string) []string {
	if slice == nil || len(slice) == 0 {
		return []string{}
	}
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func insertRules(rules map[string]rule) error {
	data := make([]rule, 0)
	parentMap := make(map[string]string)
	for _, v := range rules {
		// menu apis 去重
		v.Apis = unique(v.Apis)
		v.Menus = unique(v.Menus)
		v.FrontIds = unique(v.FrontIds)
		if v.ModuleType == 0 {
			v.ModuleType = 1
		}
		data = append(data, v)
		if _, ok := parentMap[v.ParentId]; !ok && v.ParentId != "" {
			data = append(data, rule{
				Id:         v.ParentId,
				ParentId:   "",
				Apis:       []string{},
				Menus:      []string{},
				FrontIds:   []string{},
				ListOrder:  0,
				ModuleType: v.ModuleType,
				AuthType:   v.AuthType,
				Status:     1,
			})
			parentMap[v.ParentId] = v.ParentId
		}
	}
	_, err := dao.SysRules.Ctx(context.TODO()).Data(data).Insert()
	if err != nil {
		return err
	}
	return nil
}

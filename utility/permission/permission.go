package permission

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wujchao/sys-base/consts"
	"sync"
)

var permissionApis = sync.Map{}

func GetPermissionKey(uid string) string {
	return consts.PermissionCacheKey + uid
}

func SetPermissionForUid(uid string, apis []string) error {
	if apis == nil {
		apis = []string{}
	}
	if len(apis) == 0 {
		_, err := g.Redis().Del(context.TODO(), GetPermissionKey(uid))
		permissionApis.Delete(uid)
		return err
	}
	apiMap := make(map[string]interface{})
	for _, api := range apis {
		apiMap[api] = true
	}
	_, err := g.Redis().Set(context.TODO(), GetPermissionKey(uid), apiMap)
	permissionApis.Store(uid, apiMap)
	return err
}

// CheckPermission 检查用户是否有权限访问
func CheckPermission(uid string, url string) bool {
	g.Log().Debug(context.TODO(), "CheckPermission", "uid", uid, "url", url)
	apiMap, ok := permissionApis.Load(uid)
	if !ok {
		// 缓存未命中, 从Redis中获取
		redisApiMap, err := g.Redis().Get(context.TODO(), GetPermissionKey(uid))
		if err != nil {
			return false
		}
		apiMap = redisApiMap.Map()
		permissionApis.Store(uid, apiMap)
	}
	if _, ok := apiMap.(map[string]interface{}); !ok {
		return false
	}

	return true
}

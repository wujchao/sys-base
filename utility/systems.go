package utility

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
	"sync"
	"sys-base/consts"
	"sys-base/internal/model"
)

func EncodeRuleApi(item ghttp.RouterItem) string {
	re := regexp.MustCompile(consts.ApiVersionRegex)
	url := re.ReplaceAllString(item.Route, "")
	return item.Method + ":" + url
}

func DecodeRuleApi(ruleApi string) (method, url string, err error) {
	if ruleApi == "" {
		return "", "", errors.New("ruleApi is empty")
	}
	data := strings.SplitN(ruleApi, ":", 2)
	if len(data) < 2 {
		return "", "", errors.New("ruleApi is error")
	}
	return data[0], data[1], nil
}

// EncryptPassword GeneratePassword 密码加密
func EncryptPassword(pass, slat string) string {
	costs := 10
	if slat != "" {
		costs = len(slat)
	}
	b, _ := PasswordHash([]byte(pass), costs)
	return string(b)
}

// PasswordHash 创建密码的散列值;costs为算法的cost,范围4~31,默认10;注意:值越大越耗时.
func PasswordHash(password []byte, costs ...int) ([]byte, error) {
	var cost int
	if len(costs) == 0 {
		cost = 10
	} else {
		cost = costs[0]
		if cost < 4 {
			cost = 4
		} else if cost > 31 {
			cost = 15
		}
	}

	res, err := bcrypt.GenerateFromPassword(password, cost)
	return res, err
}

// PasswordVerifyKgo PasswordVerify 验证密码是否和散列值匹配.
func PasswordVerifyKgo(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// PasswordVerify 验证密码是否和散列值匹配
func PasswordVerify(password, hash []byte) bool {
	return PasswordVerifyKgo(password, hash)
}

// PermissionApis 用户权限缓存列表
var PermissionApis = sync.Map{}

// GetPermissionKey 用户权限缓存列表 Key
func GetPermissionKey(uid string) string {
	return consts.PermissionCacheKey + gconv.String(uid)
}

// SetPermissionForUid 设置用户权限缓存
func SetPermissionForUid(uid string, apis []string) error {
	if apis == nil {
		apis = []string{}
	}
	if len(apis) == 0 {
		_, err := g.Redis().Del(context.TODO(), GetPermissionKey(uid))
		PermissionApis.Delete(uid)
		return err
	}
	apiMap := make(map[string]interface{})
	for _, api := range apis {
		apiMap[api] = true
	}
	_, err := g.Redis().Set(context.TODO(), GetPermissionKey(uid), apiMap)
	PermissionApis.Store(uid, apiMap)
	return err
}

// CheckPermission 检查用户是否有权限访问
func CheckPermission(uid string, url string) bool {
	apiMap, ok := PermissionApis.Load(uid)
	if !ok {
		// 缓存未命中, 从Redis中获取
		redisApiMap, err := g.Redis().Get(context.TODO(), GetPermissionKey(uid))
		if err != nil {
			return false
		}
		apiMap = redisApiMap.Map()
		PermissionApis.Store(uid, apiMap)
	}
	if _, ok := apiMap.(map[string]interface{}); !ok {
		return false
	}

	return true
}

// StatusToBool 1,2 表示的status转为bool
func StatusToBool[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](status T) bool {
	if status == 1 {
		return true
	}
	return false
}

// IsSystemUser 是否系统端管理员
func IsSystemUser(user *model.ContextUser) bool {
	if (user.UserTypes == consts.UserTypeAdmin || user.UserTypes == consts.UserTypeSystem) && user.OrgId != "" {
		return true
	}
	return false
}

// IsPermissionsData 数据权限校验
func IsPermissionsData(user *model.ContextUser, data any) bool {
	dataMap := gconv.Map(data)
	// 系统端管理员
	if IsSystemUser(user) {
		return true
	}
	// 仅可以操作自身数据
	if orgId := dataMap["orgId"]; gconv.String(orgId) != user.OrgId {
		return false
	}
	return true
}

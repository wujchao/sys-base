package gcasbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	adapter "github.com/hailaz/gf-casbin-adapter/v2"
	"sync"
)

type sCasbin struct {
	Enforcer *casbin.Enforcer
}

var singletonCasbin *sCasbin
var once sync.Once

func GetCasbin() *sCasbin {
	if singletonCasbin == nil {
		once.Do(func() {
			e := initCasbin()
			singletonCasbin = &sCasbin{
				Enforcer: e,
			}
		})
	}
	return singletonCasbin
}

func initCasbin() *casbin.Enforcer {
	var err error
	casbinDb, err := gdb.New(gdb.ConfigNode{
		Type: "mysql",
		Link: g.Cfg().MustGet(context.TODO(), "database.default.link").String(),
	})
	if err != nil {
		panic(err)
	}
	a := adapter.NewAdapter(adapter.Options{GDB: casbinDb})
	modelConfPath := gfile.Join(gfile.Pwd(), "rbac_model.conf")
	g.Dump(modelConfPath)
	Enforcer, err := casbin.NewEnforcer(modelConfPath, a)
	if err != nil {
		panic(err)
	}
	err = Enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	return Enforcer
}

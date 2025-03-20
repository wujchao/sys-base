package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

func init() {
	cache = gcache.New()
	ctx := context.Background()
	redisCacheAddress := g.Cfg().MustGet(ctx, "redis.cache.address", "127.0.0.1:6379")
	redisCacheDb := g.Cfg().MustGet(ctx, "redis.cache.db", 9)
	redis, err := gredis.New(&gredis.Config{
		Address: redisCacheAddress.String(),
		Db:      redisCacheDb.Int(),
	})
	if err != nil {
		panic(err)
	}
	cache.SetAdapter(gcache.NewAdapterRedis(redis))
	fmt.Println("Cache start success")
}

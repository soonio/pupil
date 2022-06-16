package logic

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/soonio/pupil/app"
)

type cacheLogic struct{}

var Cache = new(cacheLogic)

const CachePrefix = "c:" // 缓存前缀

// Get 获取
func (t *cacheLogic) Get(key string) (string, error) {
	return app.Redis.Get(context.Background(), CachePrefix+key).Result()
}

// Has 判断key是否存在
func (t *cacheLogic) Has(key string) (bool, error) {
	c, err := app.Redis.Exists(context.TODO(), CachePrefix+key).Result()
	return c > 0, err // c = 1
}

// Set 设置
func (t *cacheLogic) Set(key string, value string, duration time.Duration) error {
	_, err := app.Redis.Set(context.Background(), CachePrefix+key, value, duration).Result()
	return err
}

// Del 删除缓存
func (t *cacheLogic) Del(key string) error {
	return app.Redis.Del(context.Background(), CachePrefix+key).Err()
}

// Remember 记住 TODO 并发产生的缓存穿透
func (t *cacheLogic) Remember(key string, fn func(params any) string, duration time.Duration, params any) (string, error) {
	str, err := app.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		value := fn(params)
		err = t.Set(key, value, duration)
		return value, err
	}
	return str, err
}

// Forget 忘记
func (t *cacheLogic) Forget(key string) error {
	return t.Del(key)
}

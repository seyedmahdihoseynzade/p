package main

var redis = `package {redisDS}

import (
	"code.ts.co.ir/gaplib/golib/cache/redis"
    "sync"
    "time"
)

type RedisDataSource struct {
	redis.BaseCache
}

var (
	redisDS *RedisDataSource
	doOnce  sync.Once
)

func initRedisDS() {
	redisDS = &RedisDataSource{
		BaseCache:    redis.NewCache("{lowerModelName}Cache", "{lowerModelName}_", time.Hour*12),
	}
}

func GetRedisDS() *RedisDataSource {
	doOnce.Do(initRedisDS)
	return redisDS
}
`

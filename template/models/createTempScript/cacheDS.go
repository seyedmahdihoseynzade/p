package main

var cacheDS = `package {dataSources}

import (
	"apiGolang/models/{lowerModelName}/{dataSources}/{redisDS}"
)

type {upperModelName}CacheDS interface {
}

func GetCacheDS() {upperModelName}CacheDS {
	return redisDS.GetRedisDS()
}



`

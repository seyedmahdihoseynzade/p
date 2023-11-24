package cacheDS

/*
type RedisDataSource struct {
	redis.BaseCache
}

var (
	redisDS *RedisDataSource
	doOnce  sync.Once
)

func initRedisDS() {
	redisDS = &RedisDataSource{
		BaseCache: redis.NewCache("modelNameCache", "modelName_", like time.Hour * 12),
	}
}

func GetRedisDS() *RedisDataSource {
	doOnce.Do(initRedisDS)
	return redisDS
}

func (redisDS *RedisDataSource) GetFromCacheByID(spanCtx context.Context, groupID int64) (model dataModel.Group, err error) {
	// cache methods doesnt need logger
	_, err = redisDS.Get(spanCtx, redisDS.PrepareKey(redisDS.DefaultCachePrefix, groupID), &model)
	return model, err
}
*/

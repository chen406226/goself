package initialize

import (
	"github.com/go-redis/redis"
	"student/global"
)

func Redis()  {
	redisCfg := global.GL_CONFIG.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		global.GL_LOG.Error(err)
	} else {
		global.GL_LOG.Info("redis connect ping response:", pong)
		global.GL_REDIS = rdb
	}
}
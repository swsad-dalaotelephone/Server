package middleware

import (
	"baobaozhuan/config"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(
		config.RedisConfig.Host+":"+strconv.Itoa(config.RedisConfig.Port),
		config.RedisConfig.Password,
		time.Minute)
	return cache.Cache(&cacheStore)
}

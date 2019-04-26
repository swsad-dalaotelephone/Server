package sessionMiddleware

import (
	"baobaozhuan/config"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// register session named name
func RegisterSession(name string) gin.HandlerFunc {
	store, _ := redis.NewStore(
		10,
		"tcp",
		config.RedisConfig.Host+":"+strconv.Itoa(config.RedisConfig.Port),
		config.RedisConfig.Password,
		[]byte(config.RedisConfig.SessionKey))
	return sessions.Sessions(name, store)
}

package auth

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "invalid session token",
			})
			c.Error(errors.New("invalid session token"))
			log.ErrorLog.Println("invalid session token")
		} else {
			id := userId.(string)
			users, _ := userModel.GetUsersByStrKey("id", id)
			if len(users) == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"msg": "invalid session token",
				})
				c.Error(errors.New("invalid session token"))
				log.ErrorLog.Println("invalid session token")
			} else {
				c.Set("user", users[0])
			}
			c.Next()
		}
	}
}

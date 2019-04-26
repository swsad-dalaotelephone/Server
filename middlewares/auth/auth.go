package auth

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Invalid session token",
			})
			c.Error(errors.New("Invalid session toke"))
		} else {
			c.Next()
		}
	}
}

package userController

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
user logout
*/
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid session token",
		})
		c.Error(errors.New("Invalid session token"))
		return
	}
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully logout",
	})
}

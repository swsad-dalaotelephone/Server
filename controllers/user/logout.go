package userController

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

/*
user logout
*/
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid session token",
		})
		log.ErrorLog.Println("invalid session token")
		c.Error(errors.New("invalid session token"))
		return
	}
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully logout",
	})
	log.InfoLog.Println("successfully logout")
}

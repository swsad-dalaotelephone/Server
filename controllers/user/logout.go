package userController

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

/*
Logout : log out
require: cookie
return: msg
*/
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully logout",
	})
	log.InfoLog.Println("successfully logout")
}

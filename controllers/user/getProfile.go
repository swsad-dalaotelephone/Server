package userController

import (
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetProile : get user profile
require: cookie
return: user
*/
func GetProfile(c *gin.Context) {

	user := c.MustGet("user").(userModel.User)

	userJson, err := util.StructToJsonStr(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "successfully register",
		"user": userJson,
	})
	log.InfoLog.Println("successfully register")

}

package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
UpdateProfile : modify password
require: user profile body
return: msg
*/
func UpdateProfile(c *gin.Context) {

	var newUser userModel.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	oldUsers, err := userModel.GetUsersByStrKey("id", newUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	if len(oldUsers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "can not find user",
		})
		log.ErrorLog.Println("can not find user")
		c.Error(errors.New("can not find user"))
		return
	}

	// todo: field check

	if err := userModel.UpdateUser(newUser); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "successfully update profile",
		})
		log.InfoLog.Println("successfully update profile")
		c.Error(errors.New("successfully update profile"))
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to update profile",
		})
		log.ErrorLog.Println("fail to update profile")
		c.Error(errors.New("fail to update profile"))
	}
}
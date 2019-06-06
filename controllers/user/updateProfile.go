package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/task"
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

	user := c.MustGet("user").(userModel.User)
	id := user.Id

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
	newUser.Id = id

	oldUsers, err := userModel.GetUsersByStrKey("id", id)
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

	// user auth to check user
	oldUser := c.MustGet("user").(userModel.User)
	if oldUser.Id != id {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "permission denied",
			})
			log.ErrorLog.Println("permission denied")
			c.Error(errors.New("permission denied"))
			return
		}
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

	// modify "AccepterName" in all acceptance of this user
	if newUser.NickName != oldUsers[0].NickName {
		acceptances, err := taskModel.GetAcceptancesByStrKey("accepter_id", oldUsers[0].Id)
		if err != nil {
			log.ErrorLog.Println(err)
		} else {
			for _, item := range acceptances {
				item.AccepterName = newUser.NickName
				taskModel.UpdateAcceptance(item)
			}
		}
	}
}

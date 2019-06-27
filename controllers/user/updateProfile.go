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
require: user profile body, cookie
return: msg
*/
func UpdateProfile(c *gin.Context) {

	user := c.MustGet("user").(userModel.User)

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
	newUser.Id = user.Id
	newUser.Password = user.Password
	newUser.Phone = user.Phone

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
	if newUser.NickName != user.NickName {
		acceptances, err := taskModel.GetAcceptancesByStrKey("accepter_id", user.Id)
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

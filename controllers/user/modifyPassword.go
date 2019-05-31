package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
ModifyPassword : modify password
require: id, old_pass, new_pass
return: msg
*/
func ModifyPassword(c *gin.Context) {

	id := c.PostForm("id")
	old_pass := c.PostForm("old_pass")
	new_pass := c.PostForm("new_pass")

	// check user exist or not
	users, err := userModel.GetUsersByStrKey("id", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "user is not exist",
		})
		log.ErrorLog.Println("user is not exist")
		c.Error(errors.New("user is not exist"))
		return
	}

	// check old_pass correct or not
	if util.MD5(old_pass) != users[0].Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "old password is incorrect",
		})
		log.ErrorLog.Println("old password is incorrect")
		c.Error(errors.New("old password is incorrect"))
		return
	}

	users[0].Password = util.MD5(new_pass)
	if err := userModel.UpdateUser(users[0]); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "successfully modify password",
		})
		log.ErrorLog.Println("successfully modify password")
		c.Error(errors.New("successfully modify password"))
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
	}

}

package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
GetPreferencesById : get all tag id of user preferences name by user id
require: user id
return: tag ids list
*/
func GetPreferencesById(c *gin.Context) {

	// get user id
	// id := c.Query("user_id")
	user := c.MustGet("user").(userModel.User)
	id := user.Id

	preferences, err := userModel.GetPreferencesByStrKey("user_id", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch preference list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(preferences) > 0 {
		tagIds := make([]int, len(preferences))
		for i := 0; i < len(preferences); i++ {
			tagIds[i] = preferences[i].TagId
		}
		c.JSON(http.StatusOK, gin.H{
			"tagIds": tagIds,
		})
		log.InfoLog.Println(id, len(preferences), "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id",
		})
		log.ErrorLog.Println("invalid id")
		c.Error(errors.New("invalid id"))
	}
}

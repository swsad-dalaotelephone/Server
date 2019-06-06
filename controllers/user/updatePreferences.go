package userController

import (
	"net/http"
	"strconv"

	"github.com/swsad-dalaotelephone/Server/models/preference"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
UpdatePreferences : get all tag id of user preferences name by user id
require: user id
return: tag ids list
*/
func UpdatePreferences(c *gin.Context) {

	user := c.MustGet("user").(userModel.User)
	preferences := c.PostFormArray("preferences")

	oldPreferences, err := preferenceModel.GetPreferencesByStrKey("user_id", user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(oldPreferences) > 0 {
		// delete preferences
		for _, p := range oldPreferences {
			err := preferenceModel.DeletePreferenceById(p.Id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": err.Error(),
				})
				log.ErrorLog.Println(err)
				c.Error(err)
				return
			}
		}
	}

	// add new preferences
	for _, tag := range preferences {
		var preference preferenceModel.Preference
		preference.UserId = user.Id
		preference.TagId, _ = strconv.Atoi(tag)
		preference, _ = preferenceModel.AddPreference(preference)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	log.InfoLog.Println("success")

}

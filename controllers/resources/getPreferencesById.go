package resourcesController

import (
	"github.com/gin-gonic/gin"
)

/*
GetPreferencesById : get all tag id of user preferences name by user id
require: user id
return: tag ids list
*/
func GetPreferencesById(c *gin.Context) {

	// get user id
	id := c.Query("user_id")
	
	preferences, err := preferenceModel.GetPreferencesByStrKey("user_id", id)	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch preference list"
		})
		log.ErrorLog.Println(err)	
		return	
	}

	if len(preferences) > 0 {
		var tag_ids := make([]int, len(preferences))
		for i := 0; i < len(preferences); i++ {
			tag_ids[i] = preferences[i].tag_id
		}
		c.JSON(http.StatusOK, gin.H{
			"tag_ids": tag_ids
		})		
		log.InfoLog.Println(id, len(preferences), "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id"
		})
		log.ErrorLog.Println("invalid id")
	}
}

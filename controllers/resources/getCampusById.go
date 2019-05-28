package resourcesController

import (
	"github.com/swsad-dalaotelephone/Server/models/resources"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
GetCampusById : get campus name by campus id
require: campus id
return: campus name
*/
func GetCampusById(c *gin.Context) {
	
	// get campus id
	id := c.Query("campus_id")
	
	campuses, err := campusModel.GetCampusesByIntKey("campus_id", id)	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch campus list"
		})
		log.ErrorLog.Println(err)	
		return
	}

	if len(campuses) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"campus_name": campuses[0].name
		})		
		log.InfoLog.Println(id, campuses[0].name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id"
		})
		log.ErrorLog.Println("invalid id")
	}
}

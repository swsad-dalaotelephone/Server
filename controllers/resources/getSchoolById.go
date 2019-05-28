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
GetSchoolById : get school name by school id
require: school id
return: school name
*/
func GetSchoolById(c *gin.Context) {
	
	// get school id
	id := c.Query("school_id")
	
	schools, err := schoolModel.GetSchoolsByIntKey("school_id", id)	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch school list"
		})
		log.ErrorLog.Println(err)	
		return
	}

	if len(schools) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"school_name": schools[0].name
		})		
		log.InfoLog.Println(id, schools[0].name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id"
		})
		log.ErrorLog.Println("invalid id")
	}
}

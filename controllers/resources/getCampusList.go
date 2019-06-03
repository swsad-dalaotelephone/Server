package resourcesController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/campus"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetCampusList : get campus names list
require:
return: campus names list
*/
func GetCampusList(c *gin.Context) {

	// get all campus
	campuses, err := campusModel.GetAllCampuses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch campuses list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(campuses) > 0 {
		campusesJson, err := util.StructToJsonStr(campuses)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"campuses": campusesJson,
		})
		log.InfoLog.Println(len(campuses), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "campus list is empty",
		})
		log.ErrorLog.Println("campus list is empty")
		c.Error(errors.New("campus list is empty"))
	}
}

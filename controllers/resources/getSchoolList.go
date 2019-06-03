package resourcesController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/school"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetSchoolList : get school names list
require:
return: school names list
*/
func GetSchoolList(c *gin.Context) {

	// get all school
	schools, err := schoolModel.GetAllSchools()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch school list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(schools) > 0 {
		schoolsJson, err := util.StructToJsonStr(schools)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"schools": schoolsJson,
		})
		log.InfoLog.Println(len(schools), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "school list is empty",
		})
		log.ErrorLog.Println("school list is empty")
		c.Error(errors.New("school list is empty"))
	}
}

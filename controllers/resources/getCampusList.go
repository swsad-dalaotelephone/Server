package resourcesController

import (
	"github.com/gin-gonic/gin"
)

/*
GetCampusList : get campus names list
require:
return: campus names list
*/
func GetCampusList(c *gin.Context) {
	
	// get all campus	
	campuses, err := campuseModel.GetAllcampuses()	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch campuses list"
		})
		log.ErrorLog.Println(err)	
		return	
	}

	if len(campuses) > 0 {
		campusesJson, err := util.StructToJson(campuses)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error"
			})
			log.ErrorLog.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"campuses": campusesJson
		})		
		log.InfoLog.Println(id, len(campuses), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "campuses list is empty"
		})
		log.ErrorLog.Println("campuses list is empty")
	}
}

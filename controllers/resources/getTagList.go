package resourcesController

import (
	"github.com/gin-gonic/gin"
)

/*
GetTagList : get tag names list
require:
return: tag names list
*/
func GetTagList(c *gin.Context) {

	// get all tag	
	tags, err := schoolModel.GetAllTags()	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch tag list"
		})
		log.ErrorLog.Println(err)	
		return	
	}

	if len(tags) > 0 {
		tagsJson, err := util.StructToJson(tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error"
			})
			log.ErrorLog.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"tags": tagsJson
		})		
		log.InfoLog.Println(id, len(tags), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "tag list is empty"
		})
		log.ErrorLog.Println("tag list is empty")
	}
}

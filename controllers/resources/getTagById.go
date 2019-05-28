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
GetTagById : get tag name by tag id
require: tag id
return: tag names
*/
func GetTagById(c *gin.Context) {

	// get tag id
	id := c.Query("tag_id")
	
	tags, err := tagModel.GetTagsByIntKey("tag_id", id)	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch tag list"
		})
		log.ErrorLog.Println(err)	
		return
	}

	if len(tags) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"tag_name": tags[0].name
		})		
		log.InfoLog.Println(id, tags[0].name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id"
		})
		log.ErrorLog.Println("invalid id")
	}
}

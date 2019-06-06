package resourceController

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/swsad-dalaotelephone/Server/models/tag"
	"github.com/swsad-dalaotelephone/Server/modules/log"

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

	iid, _ := strconv.Atoi(id)
	tags, err := tagModel.GetTagsByIntKey("id", iid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch tag list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tags) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"tag_name": tags[0].Name,
		})
		log.InfoLog.Println(id, tags[0].Name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id",
		})
		log.ErrorLog.Println("invalid id")
		c.Error(errors.New("invalid id"))
	}
}

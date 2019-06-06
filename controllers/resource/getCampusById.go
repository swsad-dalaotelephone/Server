package resourceController

import (
	"errors"
	"net/http"

	"strconv"

	"github.com/swsad-dalaotelephone/Server/models/campus"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
GetCampusById : get campus name by campus id
require: campus id
return: campus name
*/
func GetCampusById(c *gin.Context) {

	// get campus id
	// id := c.Query("campus_id")
	id := c.Param("campus_id")

	iid, _ := strconv.Atoi(id)
	campuses, err := campusModel.GetCampusesByIntKey("id", iid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch campus list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(campuses) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"campus_name": campuses[0].Name,
		})
		log.InfoLog.Println(id, campuses[0].Name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id",
		})
		log.ErrorLog.Println("invalid id")
		c.Error(errors.New("invalid id"))
	}
}

package resourceController

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/swsad-dalaotelephone/Server/models/school"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
GetSchoolById : get school name by school id
require: school id
return: school name
*/
func GetSchoolById(c *gin.Context) {

	// get school id
	// id := c.Query("school_id")
	id := c.Param("school_id")

	iid, _ := strconv.Atoi(id)
	schools, err := schoolModel.GetSchoolsByIntKey("id", iid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch school list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(schools) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"school_name": schools[0].Name,
		})
		log.InfoLog.Println(id, schools[0].Name, "success")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid id",
		})
		log.ErrorLog.Println("invalid id")
		c.Error(errors.New("invalid id"))
	}
}

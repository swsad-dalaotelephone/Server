package adController

import (
	"github.com/swsad-dalaotelephone/Server/models/ad"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
GetRecommendAds : get recommend advertisement list
require:
return:
*/
func GetRecommendAds(c *gin.Context) {

	// get all ads	
	ads, err := adModel.GetAllAds()	

	if err != nil {
		log.ErrorLog.Println(err)	
		return	
	}

	if len(ads) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"ads": ads
		})		
		log.InfoLog.Println(id, len(ads), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch ad list"
		})
		log.ErrorLog.Println("can not fetch ad list")
	}
}

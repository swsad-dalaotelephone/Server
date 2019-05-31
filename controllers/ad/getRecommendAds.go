package adController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/ad"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
GetRecommendAds : get recommend advertisement list
require:
return:
*/
func GetRecommendAds(c *gin.Context) {

	// get all ads
	ads, err := adModel.GetAdsByStrKey("image", "png")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(ads) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"ads": ads,
		})
		log.InfoLog.Println(len(ads), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch ad list",
		})
		log.ErrorLog.Println("can not fetch ad list")
		c.Error(errors.New("can not fetch ad list"))
	}
}

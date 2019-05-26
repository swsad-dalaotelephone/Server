package main

import (
	"strconv"

	"github.com/swsad-dalaotelephone/Server/config"
	. "github.com/swsad-dalaotelephone/Server/router"

	"github.com/gin-gonic/gin"
)

func init() {

}
func main() {
	// set gin debug mode
	if config.ServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := InitRouter()
	router.Run(":" + strconv.Itoa(config.ServerConfig.Port))
}

package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/swsad-dalaotelephone/Server/config"
	. "github.com/swsad-dalaotelephone/Server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// set gin debug mode
	if config.ServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := InitRouter()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

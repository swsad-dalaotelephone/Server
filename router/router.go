package router

import (
	"baobaozhuan/config"
	"baobaozhuan/controllers/user"
	"baobaozhuan/middlewares/auth"
	"baobaozhuan/middlewares/session"
	"os"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//set routers
func InitRouter() *gin.Engine {
	router := gin.Default()
	// add logger middleware
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	infoFile, _ := os.OpenFile("storage/logs/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        infoFile,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		},
	)

	subLog := zerolog.New(infoFile).With().Timestamp().Logger()
	router.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
	}))
	// set recovery middleware
	router.Use(gin.Recovery())
	// add session middleware
	router.Use(sessionMiddleware.RegisterSession(config.CookieConfig.Name))
	router.GET("/weAppLogin", userController.WeAppLogin)
	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)
	router.GET("/logout", auth.AuthMiddleware(), userController.Logout)

	return router
}

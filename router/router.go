package router

import (
	"os"
	"time"

	"github.com/swsad-dalaotelephone/Server/config"
	"github.com/swsad-dalaotelephone/Server/middlewares/auth"

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
	dname := os.Getenv("GOPATH")
	filePath := dname + "src/github.com/swsad-dalaotelephone/Server/logs/api.log"
	infoFile, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

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
	router.GET("/", func(c *gin.Context) {
		c.String(200, "welcome to baobaozhuan!")
	})

	router.Use(sessionMiddleware.RegisterSession(config.CookieConfig.Name))

	// user api
	userGroup := router.Group("/user")
	{
		userGroup.GET("/loginWeApp", userController.WeAppLogin)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/register", userController.Register)
		userGroup.GET("/logout", auth.AuthMiddleware(), userController.Logout)
		userGroup.POST("/updateProfile", auth.AuthMiddleware(), userController.Logout)
		userGroup.POST("/modifyPassword", auth.AuthMiddleware(), userController.Logout)
	}

	// task api
	taskGroup := router.Group("/task")
	{
		taskGroup.POST("/publishTask", taskController.PublishTask)
		taskGroup.GET("/stopTask", auth.AuthMiddleware(), taskController.StopTask)
		taskGroup.POST("/verifyTask", auth.AuthMiddleware(), taskController.VerifyTask)
		taskGroup.POST("/updateTask", auth.AuthMiddleware(), taskController.UpdateTask)
		taskGroup.POST("/submitTask", auth.AuthMiddleware(), taskController.SubmitTask)
		taskGroup.GET("/acceptTask", taskController.AcceptTask)
		taskGroup.GET("/getRecommendTasks", taskController.GetRecommendTasks)
		taskGroup.GET("/getTaskDetail", taskController.GetTaskDetail)
		taskGroup.GET("/quitTask", auth.AuthMiddleware(), taskController.QuitTask)
	}

	// ad api
	adGroup := router.Group("/ad")
	{
		adGroup.GET("/getRecommendAds", auth.AuthMiddleware(), adController.GetRecommendAds)
	}

	// resources api
	resourcesGroup := router.Group("resources")
	{
		resourcesGroup.GET("/getSchoolById", auth.AuthMiddleware(), resourcesController.GetSchoolById)
		resourcesGroup.GET("/getCampusById", auth.AuthMiddleware(), resourcesController.GetCampusById)
		resourcesGroup.GET("/getPreferencesById", auth.AuthMiddleware(), resourcesController.GetPreferencesById)
		resourcesGroup.GET("/getTagById", auth.AuthMiddleware(), resourcesController.GetTagById)
	}

	return router
}

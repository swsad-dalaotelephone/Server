package router

import (
	"github.com/swsad-dalaotelephone/Server/config"
	"github.com/swsad-dalaotelephone/Server/controllers/ad"
	"github.com/swsad-dalaotelephone/Server/controllers/resource"
	"github.com/swsad-dalaotelephone/Server/controllers/task"
	"github.com/swsad-dalaotelephone/Server/controllers/user"
	"github.com/swsad-dalaotelephone/Server/middlewares/auth"
	"github.com/swsad-dalaotelephone/Server/middlewares/logger"
	"github.com/swsad-dalaotelephone/Server/middlewares/session"

	"github.com/gin-gonic/gin"
)

//set routers
func InitRouter() *gin.Engine {
	router := gin.Default()
	// add logger middleware
	router.Use(logger.Logger())
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
		userGroup.GET("/weApp", userController.WeAppLogin)
		userGroup.POST("", userController.Register)
		userGroup.POST("/session", userController.Login)
		userGroup.DELETE("/session", auth.AuthMiddleware(), userController.Logout)
		userGroup.GET("/profile", auth.AuthMiddleware(), userController.GetProfile)
		//userGroup.PUT("/profile", auth.AuthMiddleware(), userController.UpdateProfile)
		userGroup.POST("/profile", auth.AuthMiddleware(), userController.UpdateProfile)
		//userGroup.PUT("/password", auth.AuthMiddleware(), userController.ModifyPassword)
		userGroup.POST("/password", auth.AuthMiddleware(), userController.ModifyPassword)
		userGroup.GET("/preferences", userController.GetPreferencesById)
		userGroup.POST("/preferences", userController.UpdatePreferences)
		//userGroup.PUT("/preferences", userController.UpdatePreferences)
		userGroup.GET("/recommendedTasks", userController.GetRecommendTasks)
		userGroup.GET("/publishedTasks", auth.AuthMiddleware(), userController.GetPublishedTasks)
		userGroup.GET("/acceptedTasks", auth.AuthMiddleware(), userController.GetAcceptedTasks)
	}

	// task api
	taskGroup := router.Group("/task")
	{
		taskGroup.POST("", auth.AuthMiddleware(), taskController.PublishTask)
		taskGroup.PUT("", auth.AuthMiddleware(), taskController.UpdateTask)
		taskGroup.GET("/:task_id", taskController.GetTaskDetail)
		taskGroup.GET("/:task_id/submittedTasks", auth.AuthMiddleware(), taskController.GetSubmittedTasks)
		//taskGroup.PUT("/:task_id/status", auth.AuthMiddleware(), taskController.StopTask)
		taskGroup.POST("/:task_id/status", auth.AuthMiddleware(), taskController.StopTask)
		taskGroup.GET("/:task_id/statistic", auth.AuthMiddleware(), taskController.GetStatistics)
		taskGroup.GET("/:task_id/statistic/downloadLink", auth.AuthMiddleware(), taskController.DownloadStatistics)
		taskGroup.POST("/:task_id/acceptance", auth.AuthMiddleware(), taskController.AcceptTask)
		taskGroup.DELETE("/:task_id/acceptance", auth.AuthMiddleware(), taskController.QuitTask)
		//taskGroup.PUT("/:task_id/acceptance/answer", auth.AuthMiddleware(), taskController.SubmitTask)
		//taskGroup.PUT("/:task_id/acceptance/result", auth.AuthMiddleware(), taskController.VerifyTask)
		taskGroup.POST("/:task_id/acceptance/answer", auth.AuthMiddleware(), taskController.SubmitTask)
		taskGroup.POST("/:task_id/acceptance/result", auth.AuthMiddleware(), taskController.VerifyTask)
	}

	// ad api
	adGroup := router.Group("/ad")
	{
		adGroup.GET("/recommendedAds", adController.GetRecommendAds)
	}

	// resource api
	resourcesGroup := router.Group("resource")
	{
		resourcesGroup.GET("/school/:school_id", resourceController.GetSchoolById)
		resourcesGroup.GET("/campus/:campus_id", resourceController.GetCampusById)
		resourcesGroup.GET("/tag/:tag_id", resourceController.GetTagById)
		resourcesGroup.GET("/schools", resourceController.GetSchoolList)
		resourcesGroup.GET("/campuses", resourceController.GetCampusList)
		resourcesGroup.GET("/tags", resourceController.GetTagList)
	}

	return router
}

package taskController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
DownloadStatistics : download statistics
require: task_id, cookie
return: link
*/
func DownloadStatistics(c *gin.Context) {

	taskId := c.Query("task_id")
	user := c.MustGet("user").(userModel.User)

	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

	tasks, err := taskModel.GetTasksByStrKey("id", taskId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "can not find the task",
		})
		log.ErrorLog.Println("can not find the task")
		c.Error(errors.New("can not find the task"))
		return
	}

	if user.Id != tasks[0].PublisherId {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "permission denied",
		})
		log.ErrorLog.Println("permission denied")
		c.Error(errors.New("permission denied"))
		return
	}

	// todo: the statistics file path on server
	path := ""
	link := path + taskId + ".csv"

	// todo: if file is not exist, then generate the file
	exist := true
	if !exist {

	}

	c.JSON(http.StatusOK, gin.H{
		"link": link,
	})
	log.InfoLog.Println(link, "success")

}

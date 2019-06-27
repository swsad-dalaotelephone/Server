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
StopTask : stop task
require: task_id, cookie
return: msg
*/
func StopTask(c *gin.Context) {

	// taskId := c.Query("task_id")
	// publisherId := c.Query("publisher_id")
	taskId := c.Param("task_id")
	user := c.MustGet("user").(userModel.User)
	publisherId := user.Id

	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing argument",
		})
		log.ErrorLog.Println("missing arugment")
		c.Error(errors.New("missing argument"))
		return
	}

	// check task_id exist or not
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
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

	if publisherId != tasks[0].PublisherId {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "permission denied",
		})
		log.ErrorLog.Println("permission denied")
		c.Error(errors.New("permission denied"))
		return
	}

	// todo: delete acceptance by task id?
	if err := taskModel.DeleteTaskById(taskId); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
		log.InfoLog.Println("success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
}

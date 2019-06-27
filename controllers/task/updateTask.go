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
UpdateTask : update task
require: task body, cookie
return: msg
*/
func UpdateTask(c *gin.Context) {

	user := c.MustGet("user").(userModel.User)
	publisherId := user.Id

	var newTask taskModel.Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	//publisher of task is not this user
	oldTasks, err := taskModel.GetTasksByStrKey("id", newTask.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(oldTasks) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not find task",
		})
		log.ErrorLog.Println("can not find task")
		c.Error(errors.New("can not find task"))
		return
	}

	if oldTasks[0].Id != publisherId {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "permission denied",
		})
		log.ErrorLog.Println("permission denied")
		c.Error(errors.New("permission denied"))
		return
	}

	// todo: field check

	if err := taskModel.UpdateTask(newTask); err == nil {
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
	}

}

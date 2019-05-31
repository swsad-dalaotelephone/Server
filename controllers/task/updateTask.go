package taskController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
UpdateTask : update task
require: task body
return: msg
*/
func UpdateTask(c *gin.Context) {

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

	oldTasks, err := taskModel.GetTasksByStrKey("task_id", newTask.Id)
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

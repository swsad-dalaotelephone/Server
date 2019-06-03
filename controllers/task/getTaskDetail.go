package taskController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
)

/*
GetTaskDetail : get task detail
require: task_id
return: task detail
*/
func GetTaskDetail(c *gin.Context) {
	taskId := c.Query("task_id")

	// get task
	tasks, err := taskModel.GetTasksByStrKey("id", taskId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch task list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tasks) > 0 {
		tasksJson, err := util.StructToJsonStr(tasks[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"task": tasksJson,
		})
		log.InfoLog.Println(taskId, tasks[0].Name, "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task list is empty",
		})
		log.ErrorLog.Println("task list is empty")
		c.Error(errors.New("task list is empty"))
	}
}

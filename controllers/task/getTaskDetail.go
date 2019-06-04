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
		if task, err := taskModel.GetTaskContent(tasks[0]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "fail to get task content",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		} else {
			tasksJson, err := util.StructToJsonStr(task)
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
			log.InfoLog.Println(taskId, task.Name, "success")
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task does not exist",
		})
		log.ErrorLog.Println("task does not exist")
		c.Error(errors.New("task does not exist"))
	}
}

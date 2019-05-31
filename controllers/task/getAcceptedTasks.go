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
GetAcceptedTasks : get accepted task
require: accepter_id
return: accepted task list
*/
func GetAcceptedTasks(c *gin.Context) {
	accepterId := c.Query("accepter_id")

	// get accepted tasks
	tasks, err := taskModel.GetTasksByStrKey("accepter_id", accepterId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch task list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tasks) > 0 {
		tasksJson, err := util.StructToJson(tasks)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tasks": tasksJson,
		})
		log.InfoLog.Println(accepterId, len(tasks), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task list is empty",
		})
		log.ErrorLog.Println("task list is empty")
		c.Error(errors.New("task list is empty"))
	}
}

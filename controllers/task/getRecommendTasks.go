package taskController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetRecommendTasks : get recommend task
require: user_id
return: msg
*/
func GetRecommendTasks(c *gin.Context) {

	userId := c.Query("user_id")

	// get undo tasks
	tasks, err := taskModel.GetUnfinishedTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch task list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tasks) > 0 {
		tasksJson, err := util.StructToJsonStr(tasks)
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
		log.InfoLog.Println(userId, len(tasks), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task list is empty",
		})
		log.ErrorLog.Println("task list is empty")
		c.Error(errors.New("task list is empty"))
	}
}

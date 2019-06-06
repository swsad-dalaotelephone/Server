package taskController

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
)

/*
GetTaskDetail : get task detail
require: task_id
return: task detail
*/
func GetTaskDetail(c *gin.Context) {
	// taskId := c.Query("task_id")
	taskId := c.Param("task_id")
	session := sessions.Default(c)
	id := session.Get("userId")
	userId := ""
	if id != nil {
		userId = id.(string)
		users, _ := userModel.GetUsersByStrKey("id", userId)
		if len(users) == 0 {
			userId = ""
		}
	}

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

	//tasks is empty
	if len(tasks) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task does not exist",
		})
		log.ErrorLog.Println("task does not exist")
		c.Error(errors.New("task does not exist"))
	}

	// get test content(questionnaire, recruitment, dataCollection)
	task, err := taskModel.GetTaskContent(tasks[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to get task content",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	taskJson, err := util.StructToJsonStr(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json convert error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	if userId != "" {
		if acceptance, err := taskModel.GetAcceptanceByTaskAccepterId(taskId, userId); err == nil {
			acceptanceJson, err := util.StructToJsonStr(acceptance)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "json convert error",
				})
				log.ErrorLog.Println(err)
				c.Error(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"task":       taskJson,
				"acceptance": acceptanceJson,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"task": taskJson,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"task": taskJson,
		})
	}

	log.InfoLog.Println(taskId, task.Name, "success")
}

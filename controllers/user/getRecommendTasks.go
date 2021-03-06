package userController

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetRecommendTasks : get recommend task
require: cookie or not
return: msg
*/
func GetRecommendTasks(c *gin.Context) {

	// userId := c.Query("user_id")
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

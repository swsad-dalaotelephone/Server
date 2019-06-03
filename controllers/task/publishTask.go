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
PublishTask : publish task
require: task body
return: msg
*/
func PublishTask(c *gin.Context) {

	var task taskModel.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	// todo: field check

	task, ok := taskModel.AddTask(task)
	if ok {
		taskJson, err := util.StructToJsonStr(task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"task": taskJson,
		})
		log.InfoLog.Println("success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "add task fail",
		})
		log.ErrorLog.Println("add task fail")
		c.Error(errors.New("add task fail"))
	}

}

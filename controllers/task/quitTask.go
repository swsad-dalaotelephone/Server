package taskController

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/modules/log"

	"github.com/gin-gonic/gin"
)

/*
QuitTask : quit task
require: task_id, accepter_id
return: msg
*/
func QuitTask(c *gin.Context) {
	// taskId := c.Query("task_id")
	// accepterId := c.Query("accepter_id")
	taskId := c.Param("task_id")
	session := sessions.Default(c)
	accepterId := session.Get("userId").(string)
	// userId := session.Get("userId")
	// if userId != accepterId {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "invalid user request",
	// 	})
	// 	log.ErrorLog.Println("invalid user request")
	// 	c.Error(errors.New("invalid user request"))
	// 	return
	// }

	if taskId == "" || accepterId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing argument",
		})
		log.ErrorLog.Println("missing argument")
		c.Error(errors.New("missing argument"))
		return
	}

	// check acceptance exist or not
	acceptance, err := taskModel.GetAcceptanceByTaskAccepterId(taskId, accepterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	exist := acceptance.Id != ""
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

	if err := taskModel.DeleteAcceptanceById(acceptance.Id); err == nil {
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

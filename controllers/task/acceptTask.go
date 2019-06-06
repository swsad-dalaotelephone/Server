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
AcceptTask : accept task
require: task_id, accepter_id
return: msg
*/
func AcceptTask(c *gin.Context) {
	// taskId := c.Query("task_id")
	// accepterId := c.Query("accepter_id")
	taskId := c.Param("task_id")
	user := c.MustGet("user").(userModel.User)
	accepterId := user.Id

	if taskId == "" || accepterId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing argument",
		})
		log.ErrorLog.Println("missing argument")
		c.Error(errors.New("missing argument"))
		return
	}

	// check task_id exist or not
	tasks, err := taskModel.GetTasksByStrKey("id", taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

	// check accepter_id exist or not
	// users, err := userModel.GetUsersByStrKey("id", accepterId)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": err.Error(),
	// 	})
	// 	log.ErrorLog.Println(err)
	// 	c.Error(err)
	// 	return
	// }

	// if len(users) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "invalid argument",
	// 	})
	// 	log.ErrorLog.Println("invalid argument")
	// 	c.Error(errors.New("invalid argument"))
	// 	return
	// }

	var acceptance taskModel.Acceptance
	acceptance.TaskId = taskId
	acceptance.AccepterId = accepterId
	// acceptance.AccepterName = users[0].NickName
	acceptance.AccepterName = user.NickName
	acceptance.Status = taskModel.StatusAcceptUnsubmitted

	if _, ok := taskModel.AddAcceptance(acceptance); ok {
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

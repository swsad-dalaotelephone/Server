package taskController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

/*
SubmitTask : submit task
require: task_id, answer, cookie
return: msg
*/
func SubmitTask(c *gin.Context) {

	// taskId := c.Query("task_id")
	// accepterId := c.Query("accepter_id")
	// answer := c.Query("answer")
	taskId := c.Param("task_id")
	user := c.MustGet("user").(userModel.User)
	accepterId := user.Id
	answer := c.PostForm("answer")

	if taskId == "" || accepterId == "" || answer == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing argument",
		})
		log.ErrorLog.Println("missing argument")
		c.Error(errors.New("missing argument"))
		return
	}

	// check task_id exist or not
	tasks, err := taskModel.GetTasksByStrKey("task_id", taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	// check accepter_id exist or not
	users, err := userModel.GetUsersByStrKey("user_id", accepterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	exist := len(tasks) == 1 && len(users) == 1
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

	acceptance, err := taskModel.GetAcceptanceByTaskAccepterId(taskId, accepterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	// todo check acceptance invalid or not
	if acceptance.Id != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "can not find acceptance record",
		})
		log.ErrorLog.Println("can not find acceptance record")
		c.Error(errors.New("can not find acceptance record"))
		return
	}

	acceptance.Status = 1 // status submitted
	acceptance.Answer = []byte(answer)

	if err := taskModel.UpdateAcceptance(acceptance); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
		log.InfoLog.Println(err)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

}

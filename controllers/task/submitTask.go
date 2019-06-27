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
		log.ErrorLog.Println("missing argument", " taskId = ", taskId, " accepterId = ", accepterId, " answer = ", answer)
		c.Error(errors.New("missing argument"))
		return
	}
	// check task exist and keep running
	tasks, err := taskModel.GetTasksByStrKey("id", taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	//exist
	if len(tasks) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task not exist",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	//running
	if tasks[0].Status != taskModel.StatusTaskRunning {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task isn't running",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
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
	if acceptance.Id == "" {
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

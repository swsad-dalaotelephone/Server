package userController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
)

type result struct {
	Acceptance taskModel.Acceptance `json:"acceptance"`
	Task       taskModel.Task       `json:"task"`
}

/*
GetAcceptedTasks : get accepted task
require: cookie
return: accepted task list
*/
func GetAcceptedTasks(c *gin.Context) {
	//accepterId := c.Query("accepter_id")
	user := c.MustGet("user").(userModel.User)
	accepterId := user.Id

	// get accepted acceptances
	acceptances, err := taskModel.GetAcceptancesByStrKeyWithTask("accepter_id", accepterId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch task list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(acceptances) > 0 {
		accepted := make([]result, 0)
		for _, item := range acceptances {
			item.Answer = nil
			accepted = append(accepted, result{item, item.Task})
		}

		acceptedJson, err := util.StructToJsonStr(accepted)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accepted": acceptedJson,
		})
		log.InfoLog.Println(accepterId, len(accepted), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "task list is empty",
		})
		log.ErrorLog.Println("task list is empty")
		c.Error(errors.New("task list is empty"))
	}
}

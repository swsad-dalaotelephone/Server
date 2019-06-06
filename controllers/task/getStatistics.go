package taskController

import (
	"errors"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

type resultItem struct {
	optionCount []int    `json:"option_count"`
	optionName  []string `json:"option_name"`
	question    string   `json:"question"`
}

type result struct {
	statistics []resultItem
}

/*
GetStatistics : get statistics of the questionnaire
require: task_id, cookie
return: link
*/
func GetStatistics(c *gin.Context) {

	// taskId := c.Query("task_id")
	taskId := c.Param("task_id")
	user := c.MustGet("user").(userModel.User)

	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid argument",
		})
		log.ErrorLog.Println("invalid argument")
		c.Error(errors.New("invalid argument"))
		return
	}

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
			"msg": "can not find the task",
		})
		log.ErrorLog.Println("can not find the task")
		c.Error(errors.New("can not find the task"))
		return
	}

	if user.Id != tasks[0].PublisherId {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "permission denied",
		})
		log.ErrorLog.Println("permission denied")
		c.Error(errors.New("permission denied"))
		return
	}

	if tasks[0].Type != "q" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid task type",
		})
		log.ErrorLog.Println("invalid task type")
		c.Error(errors.New("invalid task type"))
		return
	}

	acceptances, err := taskModel.GetAcceptancesByStrKey("task_id", taskId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(acceptances) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "can not find acceptance",
		})
		log.ErrorLog.Println("can not find acceptance")
		c.Error(errors.New("can not find acceptance"))
		return
	}

	// === init res ===

	tempAnswers, err := simplejson.NewJson(acceptances[0].Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json pasre error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	tempAnswersArr, err := tempAnswers.Get("answer").Array()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json parse error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	var res result
	for i := range tempAnswersArr {
		tempAnswer := tempAnswers.Get("answer").GetIndex(i)
		options, err := tempAnswer.Get("option").Array()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json parse error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}

		if tempAnswer.Get("type").MustString() == "m" {
			var item resultItem
			item.optionCount = make([]int, len(options))
			item.optionName = make([]string, len(options))
			item.question = "question"
			res.statistics = append(res.statistics, item)
		}
	}

	// === init res ===

	for i := 0; i < len(acceptances); i++ {
		answers, err := simplejson.NewJson(acceptances[i].Answer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json parse error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}

		answersArr, err := answers.Get("answer").Array()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json parse error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}

		for i := range answersArr {
			answer := answers.Get("answer").GetIndex(i)
			if answer.Get("type").MustString() == "m" {
				options, err := answer.Get("option").Array()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "json parse error",
					})
					log.ErrorLog.Println(err)
					c.Error(err)
					return
				}

				for j := range options {
					res.statistics[i].optionCount[answer.Get("option").GetIndex(j).MustInt()]++
				}
			}
		}
	}

	resJson, err := util.StructToJson(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json convert error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statistics": resJson,
	})
	log.InfoLog.Println(len(res.statistics), "success")

}

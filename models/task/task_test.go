package taskModel

import (
	"github.com/swsad-dalaotelephone/Server/modules/util"
	"encoding/json"
	"testing"
)

func TestAddTask(t *testing.T) {
	type Img struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	img := &Img{"golang.png", "https://www.domain.com/avatar.png"}
	icon, _ := json.Marshal(img)

	res := AddTask(Task{
		PublisherId:   "xxx",
		Type:          "q",
		Name:          "dalkdj",
		BriefInfo:     "testteatlakjflksajf",
		Contract:      "123123",
		Requirements:  icon,
		RequiredCount: 1})
	t.Log(res)
}

func TestGetTask(t *testing.T) {
	tasks1, err1 := GetTasksByStrKey("id", "q")
	t.Log(tasks1, err1)
	tasks2, err2 := GetTasksByStrKey("type", "q")
	t.Log(tasks2[0], err2)
	mapData := util.JsonToMap(tasks2[0].Requirements)
	t.Log(mapData)
}

func TestGetUnfinishedTask(t *testing.T) {
	tasks1, _ := GetUnfinishedTask()
	t.Log(len(tasks1))
}

func TestDeleteTask(t *testing.T) {
	tasks, _ := GetTasksByStrKey("type", "q")
	t.Log(len(tasks))
	DeleteTaskById(tasks[0].Id)
	tasks, _ = GetTasksByStrKey("type", "q")
	t.Log(len(tasks))
}

func TestQuestionnaire(t *testing.T) {
	tasks, _ := GetTasksByIntKey("type", 1)
	for i := 0; i < len(tasks); i++ {
		DeleteTaskById(tasks[i].Id)
	}

	type Img struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	img := &Img{"golang.png", "https://www.domain.com/avatar.png"}
	icon, _ := json.Marshal(img)

	res := AddTask(Task{
		PublisherId:   "qqq",
		Type:          "q",
		Name:          "问卷test",
		BriefInfo:     "testteatlakjflksajf",
		Contract:      "123123",
		RequiredCount: 1,
		Questionnaire: Questionnaire{
			Description: "dddddd",
			Questions:   icon}})
	t.Log(res)

	tasks, _ = GetTasksByStrKey("type", "q")
	task := tasks[0]
	task, _ = GetTaskDetail(task)
	t.Log(task)
}

func TestDataCollection(t *testing.T) {

}

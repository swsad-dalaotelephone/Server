package taskModel

import (
	"encoding/json"
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/models/common"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/jinzhu/gorm"
)

// table name
const (
	TaskTableName = "task"
)

type Task struct {
	Id            string           `gorm:"column:id; type:varchar(36); primary_key; not null" json:"id"`
	PublisherId   string           `gorm:"column:publisher_id; type:varchar(36); not null; index:publisher_id_idx" json:"publisher_id"`
	Type          string           `gorm:"column:type; index:type_idx" json:"type"`
	Name          string           `gorm:"column:name" json:"name"`
	BriefInfo     string           `gorm:"column:brief_info" json:"brief_info"`
	Contact       string           `gorm:"column:contact" json:"contact"`
	Requirements  commonModel.JSON `gorm:"column:requirements" sql:"type:json" json:"requirements"`
	DDL           time.Time        `gorm:"column:ddl" json:"ddl"`
	Reward        int              `gorm:"column:reward; default:0" json:"reward"`
	TagId         int              `gorm:"column:tag_id; default:0; index:tag_id_idx" json:"tag_id"`
	RequiredCount int              `gorm:"column:required_count; default:0" json:"required_count"`
	Status        int              `gorm:"column:status; default:0" json:"status"`
	Content       commonModel.JSON `gorm:"-" json:"content"`
	CreatedAt     time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(TaskTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Task{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
	}
}

//set table name
func (u Task) TableName() string {
	return TaskTableName
}

// set task.id as uuid before create
func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	task.Id = util.GetNewUuid()
	task.DDL = time.Now()
	return nil
}

/*
 add new task.
 it will fail, if content can't be stored in database
 @parm new task
 @return isSuccessful
*/
func AddTask(task Task) (Task, bool) {
	DB.Create(&task)
	err := SaveContent(task)
	if err != nil {
		return task, false
	}
	res := DB.NewRecord(&task) //return `false` after `task` created
	return task, !res
}

// query tasks by string key
func GetTasksByStrKey(key string, value string) (tasks []Task, err error) {
	err = DB.Where(key+" = ?", value).Find(&tasks).Error
	return tasks, err
}

// query tasks by int key
func GetTasksByIntKey(key string, value int) (tasks []Task, err error) {
	err = DB.Where(key+" = ?", value).Find(&tasks).Error
	return tasks, err
}

/*
 update task info
 must GetTaskByKey first
*/
func UpdateTask(task Task) error {
	err := DB.Save(&task).Error
	return err
}

/*
delete task by id
*/
func DeleteTaskById(id string) error {
	err := DB.Where("id = ?", id).Delete(Task{}).Error
	return err
}

/*
get unfinished tasks
*/
func GetUnfinishedTasks() (tasks []Task, err error) {
	err = DB.Where("status = 0").Find(&tasks).Error
	return tasks, err
}

/*
get Content of task
according to task type, query detail{
1: questionnaire
2: dataCollection
3: recruitment
}
*/
func GetTaskWithContentById(id string) (task Task, err error) {
	err = DB.Where("id = ?", id).First(&task).Error
	if err != nil {
		return task, err
	}
	switch task.Type {
	case "q":
		var questionnaire Questionnaire
		err = DB.Model(&task).Related(&questionnaire).Error
		if content, err := util.StructToJson(questionnaire); err != nil {
			log.ErrorLog.Println(err)
		} else {
			task.Content = content
		}
		break
	case "d":
		var dataCollection DataCollection
		err = DB.Model(&task).Related(&dataCollection).Error
		if content, err := util.StructToJson(dataCollection); err != nil {
			log.ErrorLog.Println(err)
		} else {
			task.Content = content
		}
		break
	case "r":
		var recruitment Recruitment
		err = DB.Model(&task).Related(&recruitment).Error
		if content, err := util.StructToJson(recruitment); err != nil {
			log.ErrorLog.Println(err)
		} else {
			task.Content = content
		}
		break
	}
	return task, err
}

func SaveContent(task Task) error {
	switch task.Type {
	case "q":
		var questionnaire Questionnaire
		if err := json.Unmarshal(task.Content, &questionnaire); err != nil {
			log.ErrorLog.Println(err)
			return err
		} else {
			questionnaire.TaskId = task.Id
			AddQuestionnaire(questionnaire)
		}
		break
	case "d":
		var dataCollection DataCollection
		if err := json.Unmarshal(task.Content, &dataCollection); err != nil {
			log.ErrorLog.Println(err)
			return err
		} else {
			dataCollection.TaskId = task.Id
			AddDataCollection(dataCollection)
		}
		break
	case "r":
		var recruitment Recruitment
		if err := json.Unmarshal(task.Content, &recruitment); err != nil {
			log.ErrorLog.Println(err)
			return err
		} else {
			recruitment.TaskId = task.Id
			AddRecruitment(recruitment)
		}
		break
	}
	return nil
}

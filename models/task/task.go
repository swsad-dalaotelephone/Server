package taskModel

import (
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/models/common"
	"github.com/swsad-dalaotelephone/Server/models/tag"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/jinzhu/gorm"
)

// table name
const (
	TaskTableName = "task"
)

type Task struct {
	Id             string           `gorm:"column:id; type:varchar(36); primary_key; not null" json:"id"`
	PublisherId    string           `gorm:"column:public_id; type:varchar(36); not null; index:publisher_id_idx" json:"publiser_id"`
	Type           string           `gorm:"column:type; index:type_idx" json:"type"`
	Name           string           `gorm:"column:name" json:"name"`
	BriefInfo      string           `gorm:"column:brief_info" json:"brief_info"`
	Contact        string           `gorm:"column:contact" json:"contact"`
	Requirements   commonModel.JSON `gorm:"column:requirements" sql:"type:json" json:"requirements"`
	DDL            time.Time        `gorm:"column:ddl" json:"ddl"`
	Reward         int              `gorm:"column:reward; default:0" json:"reward"`
	Tag            tagModel.Tag     `gorm:"foreignkey:TagId"`
	TagId          int              `gorm:"column:tag_id; default:0; index:tag_id_idx" json:"tag_id"`
	RequiredCount  int              `gorm:"column:required_count; default:0" json:"required_count"`
	Status         int              `gorm:"column:status; default:0" json:"status"`
	CreatedAt      time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt      time.Time        `gorm:"column:updated_at" json:"-"`
	Questionnaire  Questionnaire    `gorm:"foreignkey:TaskId"`
	DataCollection DataCollection   `gorm:"foreignkey:TaskId"`
	Recruitment    Recruitment      `gorm:"foreignkey:TaskId"`
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
 add new task
 @parm new task
 @return isSuccessful
*/
func AddTask(task Task) bool {
	DB.Create(&task)
	res := DB.NewRecord(&task) //return `false` after `task` created
	return !res
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
	err = DB.Where("submited_count + finished_count < required_count").Find(&tasks).Error
	return tasks, err
}

/*
get detail of task
according to task type, query detail{
1: questionnaire
2: dataCollection
3: recruitment
}
*/
func GetTaskDetail(task Task) (Task, error) {
	var err error
	switch task.Type {
	case "q":
		err = DB.Model(&task).Related(&task.Questionnaire).Error
	case "d":
		err = DB.Model(&task).Related(&task.DataCollection).Error
	case "r":
		err = DB.Model(&task).Related(&task.Recruitment).Error
	}
	return task, err
}

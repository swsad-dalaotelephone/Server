package taskModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/models/common"
	"baobaozhuan/models/tag"
	"baobaozhuan/modules/util"
	"time"

	"github.com/jinzhu/gorm"
)

// table name
const (
	TaskTableName = "task"
)

type Task struct {
	Id            string           `gorm:"type:varchar(36); primary_key; not null" json:"task_id"`
	PublisherId   string           `gorm:"type:varchar(36); not null; index:publisher_id_idx" json:"publiser_id"`
	Type          int              `gorm:"index:type_idx" json:"type"`
	Name          string           `json:"name"`
	BriefInfo     string           `json:"brief_info"`
	Contract      string           `json:"contract"`
	Requirements  commonModel.JSON `sql:"type:json" json:"requirements"`
	Ddl           time.Time        `json:"ddl"`
	Reward        int              `json:"reward"`
	Tag           tagModel.Tag     `gorm:"foreignkey:TagId"`
	TagId         int              `gorm:"index:tag_id_idx" json:"tag_id"`
	RequiredCount int              `json:"required_count"`
	SubmitedCount int              `gorm:"default:0" json:"submited_count"`
	FinishedCount int              `gorm:"default:0" json:"finished_count"`
	CreatedAt     time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(TaskTableName) {
		DB.CreateTable(&Task{})
	}
}

//set table name
func (u Task) TableName() string {
	return TaskTableName
}

// set task.id as uuid before create
func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	task.Id = util.GetNewUuid()
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

// query task by key
func GetTaskByKey(key string, value string) (Task, error) {
	task := Task{}
	res := DB.Where(key+" = ?", value).First(&task)
	err := res.Error
	return task, err
}

/*
 update task info
 must GetTaskByKey first
*/
func UpdateTask(task Task) {
	DB.Save(&task)
}

package taskModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/models/common"
	"time"
)

// table name
const (
	QuestionnaireTableName = "questionnaire"
)

type Questionnaire struct {
	TaskId      string           `gorm:"type:varchar(36); primary_key; not null" json:"task_id"`
	Description string           `gorm:"type:text" json:"description"`
	Questions   commonModel.JSON `sql:"type:json" json:"questions"`
	Statistics  commonModel.JSON `sql:"type:json" json:"statistics"`
	CreatedAt   time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(QuestionnaireTableName) {
		DB.CreateTable(&Questionnaire{})
	}
}

//set table name
func (u Questionnaire) TableName() string {
	return QuestionnaireTableName
}

/*
 add new questionnaire
 @parm new questionnaire
 @return isSuccessful
*/
func AddQuestionnaire(questionnaire Questionnaire) bool {
	DB.Create(&questionnaire)
	res := DB.NewRecord(&questionnaire) //return `false` after `questionnaire` created
	return !res
}

// query questionnaire by key
func GetQuestionnaireByKey(key string, value string) (Questionnaire, error) {
	questionnaire := Questionnaire{}
	res := DB.Where(key+" = ?", value).First(&questionnaire)
	err := res.Error
	return questionnaire, err
}

/*
 update questionnaire info
 must GetQuestionnaireByKey first
*/
func UpdateQuestionnaire(questionnaire Questionnaire) {
	DB.Save(&questionnaire)
}

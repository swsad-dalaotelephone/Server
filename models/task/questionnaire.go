package taskModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/models/common"
	"baobaozhuan/modules/log"
	"time"
)

// table name
const (
	QuestionnaireTableName = "questionnaire"
)

type Questionnaire struct {
	TaskId      string           `gorm:"column:task_id; type:varchar(36); primary_key; not null" json:"task_id"`
	Description string           `gorm:"column:description; type:text" json:"description"`
	Questions   commonModel.JSON `gorm:"column:questions" sql:"type:json" json:"questions"`
	Statistics  commonModel.JSON `gorm:"column:statistics" sql:"type:json" json:"statistics"`
	CreatedAt   time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(QuestionnaireTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Questionnaire{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
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

// query questionnaires by string key
func GetQuestionnairesByStrKey(key string, value string) (questionnaires []Questionnaire, err error) {
	err = DB.Where(key+" = ?", value).Find(&questionnaires).Error
	return questionnaires, err
}

/*
 update questionnaire info
 must GetQuestionnaireByKey first
*/
func UpdateQuestionnaire(questionnaire Questionnaire) error {
	err := DB.Save(&questionnaire).Error
	return err
}

/*
delete questionnaire by id
*/
func DeleteQuestionnaireById(id string) error {
	err := DB.Where("id = ?", id).Delete(Questionnaire{}).Error
	return err
}

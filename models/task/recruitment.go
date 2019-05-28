package taskModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/models/common"
	"time"
)

// table name
const (
	RecruitmentTableName = "recruitment"
)

type Recruitment struct {
	TaskId          string           `gorm:"type:varchar(36); primary_key; not null" json:"task_id"`
	Time            time.Time        `json:"time"`
	Location        string           `json:"location"`
	Description     string           `gorm:"type:text" json:"description"`
	ParticipantInfo commonModel.JSON `sql:"type:json" json:"participant_info"`
	CreatedAt       time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(RecruitmentTableName) {
		DB.CreateTable(&Recruitment{})
	}
}

//set table name
func (u Recruitment) TableName() string {
	return RecruitmentTableName
}

/*
 add new recruitment
 @parm new recruitment
 @return isSuccessful
*/
func AddRecruitment(recruitment Recruitment) bool {
	DB.Create(&recruitment)
	res := DB.NewRecord(&recruitment) //return `false` after `recruitment` created
	return !res
}

// query recruitment by key
func GetRecruitmentByKey(key string, value string) (Recruitment, error) {
	recruitment := Recruitment{}
	res := DB.Where(key+" = ?", value).First(&recruitment)
	err := res.Error
	return recruitment, err
}

/*
 update recruitment info
 must GetRecruitmentByKey first
*/
func UpdateRecruitment(recruitment Recruitment) {
	DB.Save(&recruitment)
}

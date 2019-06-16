package taskModel

import (
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/models/common"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

// table name
const (
	RecruitmentTableName = "recruitment"
)

type Recruitment struct {
	TaskId          string           `gorm:"column:task_id; type:varchar(36); primary_key; not null" json:"task_id"`
	Description     string           `gorm:"column:description; type:text" json:"recruit_des"`
	StartTime       string           `gorm:"column:start_time" json:"start_time"`
	EndTime         string           `gorm:"column:end_time" json:"end_time"`
	Location        string           `gorm:"column:location" json:"location"`
	ParticipantInfo commonModel.JSON `gorm:"column:participant_info" sql:"type:json" json:"participant_info"`
	CreatedAt       time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(RecruitmentTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Recruitment{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
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
func AddRecruitment(recruitment Recruitment) (Recruitment, bool) {
	DB.Create(&recruitment)
	res := DB.NewRecord(&recruitment) //return `false` after `recruitment` created
	return recruitment, !res
}

// query recruitments by string key
func GetRecruitmentsByStrKey(key string, value string) (recruitments []Recruitment, err error) {
	err = DB.Where(key+" = ?", value).Find(&recruitments).Error
	return recruitments, err
}

// query recruitments by int key
func GetRecruitmentsByIntKey(key string, value int) (recruitments []Recruitment, err error) {
	err = DB.Where(key+" = ?", value).Find(&recruitments).Error
	return recruitments, err
}

/*
 update recruitment info
 must GetRecruitmentByKey first
*/
func UpdateRecruitment(recruitment Recruitment) error {
	err := DB.Save(&recruitment).Error
	return err
}

/*
delete recruitment by task_id
*/
func DeleteRecruitmentByTaskId(task_id string) error {
	err := DB.Where("task_id = ?", task_id).Delete(Recruitment{}).Error
	return err
}

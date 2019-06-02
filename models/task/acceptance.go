package taskModel

import (
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/models/common"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/jinzhu/gorm"
)

// table name
const (
	AcceptanceTableName = "acceptance"
)

type Acceptance struct {
	Id           string           `gorm:"column:id; type:varchar(36); primary_key; not null" json:"acceptance_id"`
	TaskId       string           `gorm:"column:task_id; type:varchar(36); not null; unique_index:task_accepter_idx; index:task_id_idx" json:"task_id"`
	AccepterId   string           `gorm:"column:accepter_id; type:varchar(36); not null; unique_index:task_accepter_idx; index:accepter_id_idx" json:"accepter_id"`
	AccepterName string           `gorm:"column:accepter_name" json:"accepter"`
	Answer       commonModel.JSON `gorm:"column:answer" sql:"type:json" json:"answer"`
	Status       int              `gorm:"column:status; default:0" json:"status"`
	Feedback     string           `gorm:"column:feedback; type:text" json:"feedback"`
	CreatedAt    time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(AcceptanceTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Acceptance{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
	}
}

//set table name
func (u Acceptance) TableName() string {
	return AcceptanceTableName
}

// set acceptance.id as uuid before create
func (acceptance *Acceptance) BeforeCreate(scope *gorm.Scope) error {
	acceptance.Id = util.GetNewUuid()
	return nil
}

/*
 add new acceptance
 @parm new acceptance
 @return isSuccessful
*/
func AddAcceptance(acceptance Acceptance) (Acceptance, bool) {
	DB.Create(&acceptance)
	res := DB.NewRecord(&acceptance) //return `false` after `acceptance` created
	return acceptance, !res
}

// query acceptances by string key
func GetAcceptancesByStrKey(key string, value string) (acceptances []Acceptance, err error) {
	err = DB.Where(key+" = ?", value).Find(&acceptances).Error
	return acceptances, err
}

// query acceptances by int key
func GetAcceptancesByIntKey(key string, value int) (acceptances []Acceptance, err error) {
	err = DB.Where(key+" = ?", value).Find(&acceptances).Error
	return acceptances, err
}

/*
get acceptance by taskId and accepterId
*/
func GetAcceptanceByTaskAccepterId(taskId, accepterId string) (acceptance Acceptance, err error) {
	err = DB.Where(&Acceptance{TaskId: taskId, AccepterId: accepterId}).First(&acceptance).Error
	return acceptance, err
}

/*
 update acceptance info
 must GetAcceptanceByKey first
*/
func UpdateAcceptance(acceptance Acceptance) error {
	err := DB.Save(&acceptance).Error
	return err
}

/*
delete acceptance by id
*/
func DeleteAcceptanceById(id string) error {
	err := DB.Where("id = ?", id).Delete(Acceptance{}).Error
	return err
}

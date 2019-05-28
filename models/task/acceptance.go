package taskModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/models/common"
	"baobaozhuan/models/user"
	"baobaozhuan/modules/util"
	"time"

	"github.com/jinzhu/gorm"
)

// table name
const (
	AcceptanceTableName = "acceptance"
)

type Acceptance struct {
	Id           string           `gorm:"type:varchar(36); primary_key; not null" json:"acceptance_id"`
	TaskId       string           `gorm:"type:varchar(36); not null; index:task_accepter_idx" json:"task_id"`
	AccepterId   string           `gorm:"type:varchar(36); not null; index:task_accepter_idx" json:"accepter_id"`
	Accepter     userModel.User   `gorm:"foreignkey:AccepterId"`
	AccepterName string           `json:"accepter"`
	Answer       commonModel.JSON `sql:"type:json" json:"answer"`
	Status       int              `json:"status"`
	Feedback     string           `gorm:"type:text" json:"feedback"`
	CreatedAt    time.Time        `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time        `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(AcceptanceTableName) {
		DB.CreateTable(&Acceptance{})
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
func AddAcceptance(acceptance Acceptance) bool {
	DB.Create(&acceptance)
	res := DB.NewRecord(&acceptance) //return `false` after `acceptance` created
	return !res
}

// query acceptance by key
func GetAcceptanceByKey(key string, value string) (Acceptance, error) {
	acceptance := Acceptance{}
	res := DB.Where(key+" = ?", value).First(&acceptance)
	err := res.Error
	return acceptance, err
}

/*
 update acceptance info
 must GetAcceptanceByKey first
*/
func UpdateAcceptance(acceptance Acceptance) {
	DB.Save(&acceptance)
}

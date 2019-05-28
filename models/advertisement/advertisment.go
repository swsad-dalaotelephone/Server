package advertismentModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/modules/util"
	"time"

	"github.com/jinzhu/gorm"
)

// table name
const (
	AdvertismentTableName = "advertisment"
)

type Advertisment struct {
	Id        string    `gorm:"type:varchar(36); primary_key; not null" json:"id"`
	Link      string    `json:"link"`
	Image     string    `json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(AdvertismentTableName) {
		DB.CreateTable(&Advertisment{})
	}
}

//set table name
func (u Advertisment) TableName() string {
	return AdvertismentTableName
}

// set advertisment.id as uuid before create
func (advertisment *Advertisment) BeforeCreate(scope *gorm.Scope) error {
	advertisment.Id = util.GetNewUuid()
	return nil
}

/*
 add new advertisment
 @parm new advertisment
 @return isSuccessful
*/
func AddAdvertisment(advertisment Advertisment) bool {
	DB.Create(&advertisment)
	res := DB.NewRecord(&advertisment) //return `false` after `advertisment` created
	return !res
}

// query advertisment by key
func GetAdvertismentByKey(key string, value string) (Advertisment, error) {
	advertisment := Advertisment{}
	res := DB.Where(key+" = ?", value).First(&advertisment)
	err := res.Error
	return advertisment, err
}

/*
 update advertisment info
 must GetAdvertismentByKey first
*/
func UpdateAdvertisment(advertisment Advertisment) {
	DB.Save(&advertisment)
}

package advertismentModel

import (
	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/util"
	"time"

	"github.com/jinzhu/gorm"
)

// table name
const (
	AdvertismentTableName = "advertisment"
)

type Advertisment struct {
	Id        string    `gorm:"column:id; type:varchar(36); primary_key; not null" json:"id"`
	Link      string    `gorm:"column:link" json:"link"`
	Image     string    `gorm:"column:image" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(AdvertismentTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Advertisment{}}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
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

// query advertisments by string key
func GetAdvertismentByStrKey(key string, value string) (advertisments []Advertisment, err error) {
	err = DB.Where(key+" = ?", value).Find(&advertisments).Error
	return advertisment, err
}

/*
 update advertisment info
 must GetAdvertismentByKey first
*/
func UpdateAdvertisment(advertisment Advertisment) error {
	err := DB.Save(&advertisment).Error
	return err
}

/* 
delete advertisment by id
 */
func DeleteAdvertismentById(id string) error {
	err := DB.Where("id = ?", id).Delete(Advertisment{}).Error
	return err
}
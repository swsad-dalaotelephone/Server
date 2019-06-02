package adModel

import (
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/jinzhu/gorm"
)

// table name
const (
	AdTableName = "ad"
)

type Ad struct {
	Id        string    `gorm:"column:id; type:varchar(36); primary_key; not null" json:"id"`
	Link      string    `gorm:"column:link" json:"link"`
	Image     string    `gorm:"column:image" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(AdTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Ad{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
	}
}

// set table name
func (u Ad) TableName() string {
	return AdTableName
}

// set ad.id as uuid before create
func (ad *Ad) BeforeCreate(scope *gorm.Scope) error {
	ad.Id = util.GetNewUuid()
	return nil
}

/*
 add new ad
 @parm new ad
 @return isSuccessful
*/
func AddAd(ad Ad) (Ad, bool) {
	DB.Create(&ad)
	res := DB.NewRecord(&ad) //return `false` after `ad` created
	return ad, !res
}

// query ads by string key
func GetAdsByStrKey(key string, value string) (ads []Ad, err error) {
	err = DB.Where(key+" = ?", value).Find(&ads).Error
	return ads, err
}

/*
 update ad info
 must GetAdByKey first
*/
func UpdateAd(ad Ad) error {
	err := DB.Save(&ad).Error
	return err
}

/*
delete ad by id
*/
func DeleteAdById(id string) error {
	err := DB.Where("id = ?", id).Delete(Ad{}).Error
	return err
}

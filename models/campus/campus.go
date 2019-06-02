package campusModel

import (
	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

// table name
const (
	CampusTableName = "campus"
)

type Campus struct {
	Id   int    `gorm:"column:id; primary_key; not null; AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name; not null; unique" sql:"not null" json:"name"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(CampusTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Campus{}).Error; err != nil {
			log.ErrorLog.Println(err)
			return
		}
		initCampus()
	}
}

//set table name
func (u Campus) TableName() string {
	return CampusTableName
}

/*
 add new campus
 @parm new campus
 @return isSuccessful
*/
func AddCampus(campus Campus) (Campus, bool) {
	DB.Create(&campus)
	res := DB.NewRecord(&campus) //return `false` after `campus` created
	return campus, !res
}

// query campuses by string key
func GetCampusesByStrKey(key string, value string) (campuses []Campus, err error) {
	err = DB.Where(key+" = ?", value).Find(&campuses).Error
	return campuses, err
}

// query campuses by int key
func GetCampusesByIntKey(key string, value int) (campuses []Campus, err error) {
	err = DB.Where(key+" = ?", value).Find(&campuses).Error
	return campuses, err
}

/*
 update campus info
 must GetCampusByKey first
*/
func UpdateCampus(campus Campus) error {
	err := DB.Save(&campus).Error
	return err
}

/*
 get all campuses
*/
func GetAllCampuses() (campuses []Campus, err error) {
	err = DB.Find(&campuses).Error
	return campuses, err
}

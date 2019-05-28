package userModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/modules/log"
)

// table name
const (
	CampusTableName = "campus"
)

type Campus struct {
	Id   int    `gorm:"column:id; primary_key; not null; AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name; not null; unique" sql:"not null" json:"name"`
}

/*
add 5 campus of sysu
*/
func initCampus() {
	AddCampus(Campus{1, "广州东校园"})
	AddCampus(Campus{2, "广州南校园"})
	AddCampus(Campus{3, "广州北校园"})
	AddCampus(Campus{4, "珠海校区"})
	AddCampus(Campus{5, "深圳校区"})

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
func AddCampus(campus Campus) bool {
	DB.Create(&campus)
	res := DB.NewRecord(&campus) //return `false` after `campus` created
	return !res
}

// query campuses by string key
func GetCampusesByStrKey(key string, value string) (campuses []Campus, err error) {
	res := DB.Where(key+" = ?", value).First(&campuses)
	err = res.Error
	return campuses, err
}

// query campuses by int key
func GetCampusesByIntKey(key string, value int) (campuses []Campus, err error) {
	res := DB.Where(key+" = ?", value).Find(&campuses)
	err = res.Error
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

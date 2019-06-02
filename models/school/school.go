package schoolModel

import (
	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

// table name
const (
	SchoolTableName = "school"
)

type School struct {
	Id   int    `gorm:"column:id; primary_key; not null; AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name; not null; unique" sql:"not null" json:"name"`
	Type string `gorm:"column:type; not null; index:type_idx" sql:"not null" json:"type"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(SchoolTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&School{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
		initSchool()
	}
}

//set table name
func (u School) TableName() string {
	return SchoolTableName
}

/*
 add new school
 @parm new school
 @return isSuccessful
*/
func AddSchool(school School) (School, bool) {
	DB.Create(&school)
	res := DB.NewRecord(&school) //return `false` after `school` created
	return school, !res
}

// query schools by string key
func GetSchoolsByStrKey(key string, value string) (schools []School, err error) {
	err = DB.Where(key+" = ?", value).Find(&schools).Error
	return schools, err
}

// query schools by int key
func GetSchoolsByIntKey(key string, value int) (schools []School, err error) {
	err = DB.Where(key+" = ?", value).Find(&schools).Error
	return schools, err
}

/*
 update school info
 must GetSchoolByKey first
*/
func UpdateSchool(school School) error {
	err := DB.Save(&school).Error
	return err
}

/*
 get all schools
*/
func GetAllSchools() (schools []School, err error) {
	err = DB.Find(&schools).Error
	return schools, err
}

/*
delete school by id
*/
func DeleteSchoolById(id int) error {
	err := DB.Where("id = ?", id).Delete(School{}).Error
	return err
}

/*
delete school by name
*/
func DeleteSchoolByName(name string) error {
	err := DB.Where("name = ?", name).Delete(School{}).Error
	return err
}

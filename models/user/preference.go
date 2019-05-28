package userModel

import (
	. "baobaozhuan/database"
	"baobaozhuan/modules/log"
	"time"
)

// table name
const (
	PreferenceTableName = "preference"
)

type Preference struct {
	Id        int       `gorm:"column:id; primary_key; not null; AUTO_INCREMENT" json:"id"`
	UserId    string    `gorm:"column:user_id; type:varchar(36); index: user_id_idx; unique_index: user_tag_idx; not null" json:"user_id"`
	TagId     int       `gorm:"column:tag_id; index:tag_id_idx; index: tag_id_idx; unique_index: user_tag_idx; not null" json:"tag_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(PreferenceTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Preference{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
	}
}

//set table name
func (u Preference) TableName() string {
	return PreferenceTableName
}

/*
 add new preference, if exist this (preference)
 @parm new preference
 @return isSuccessful
*/
func AddPreference(preference Preference) bool {
	DB.Create(&preference)
	res := DB.NewRecord(&preference) //return `false` after `preference` created
	return !res
}

// query preferences by string key
func GetPreferencesByStrKey(key string, value string) (preferences []Preference, err error) {
	res := DB.Where(key+" = ?", value).Find(&preferences)
	err = res.Error
	return preferences, err
}

// query preferences by int key
func GetPreferencesByIntKey(key string, value int) (preferences []Preference, err error) {
	res := DB.Where(key+" = ?", value).Find(&preferences)
	err = res.Error
	return preferences, err
}

/*
 update preference info
 must GetPreferenceByKey first
*/
func UpdatePreference(preference Preference) error {
	err := DB.Save(&preference).Error
	return err
}

/*
 get all preferences
*/
func GetAllPreferences() (preferences []Preference, err error) {
	err = DB.Find(&preferences).Error
	return preferences, err
}

/*
delete preferences by id
*/
func DeletePreferenceById(id int) error {
	err := DB.Where("id = ?", id).Delete(Preference{}).Error
	return err
}

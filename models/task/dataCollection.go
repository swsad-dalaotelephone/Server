package taskModel

import (
	. "baobaozhuan/database"
	"time"
)

// table name
const (
	DataCollectionTableName = "data_collection"
)

type DataCollection struct {
	TaskId      string    `gorm:"type:varchar(36); primary_key; not null" json:"task_id"`
	Description string    `gorm:"type:text" json:"description"`
	SubmitWay   string    `json:"submit_way"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(DataCollectionTableName) {
		DB.CreateTable(&DataCollection{})
	}
}

//set table name
func (u DataCollection) TableName() string {
	return DataCollectionTableName
}

/*
 add new dataCollection
 @parm new dataCollection
 @return isSuccessful
*/
func AddDataCollection(dataCollection DataCollection) bool {
	DB.Create(&dataCollection)
	res := DB.NewRecord(&dataCollection) //return `false` after `dataCollection` created
	return !res
}

// query dataCollection by key
func GetDataCollectionByKey(key string, value string) (DataCollection, error) {
	dataCollection := DataCollection{}
	res := DB.Where(key+" = ?", value).First(&dataCollection)
	err := res.Error
	return dataCollection, err
}

/*
 update dataCollection info
 must GetDataCollectionByKey first
*/
func UpdateDataCollection(dataCollection DataCollection) {
	DB.Save(&dataCollection)
}

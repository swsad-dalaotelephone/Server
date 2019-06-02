package taskModel

import (
	"time"

	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

// table name
const (
	DataCollectionTableName = "data_collection"
)

type DataCollection struct {
	TaskId      string    `gorm:"column:task_id; type:varchar(36); primary_key; not null" json:"task_id"`
	Description string    `gorm:"column:description; type:text" json:"description"`
	SubmitWay   string    `gorm:"column:submit_way" json:"submit_way"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(DataCollectionTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DataCollection{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
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
func AddDataCollection(dataCollection DataCollection) (DataCollection, bool) {
	DB.Create(&dataCollection)
	res := DB.NewRecord(&dataCollection) //return `false` after `dataCollection` created
	return dataCollection, !res
}

// query dataCollections by string key
func GetDataCollectionsByStrKey(key string, value string) (dataCollections []DataCollection, err error) {
	err = DB.Where(key+" = ?", value).Find(&dataCollections).Error
	return dataCollections, err
}

/*
 update dataCollection info
 must GetDataCollectionByKey first
*/
func UpdateDataCollection(dataCollection DataCollection) error {
	err := DB.Save(&dataCollection).Error
	return err
}

/*
delete dataCollection by id
*/
func DeleteDataCollectionById(id string) error {
	err := DB.Where("id = ?", id).Delete(DataCollection{}).Error
	return err
}

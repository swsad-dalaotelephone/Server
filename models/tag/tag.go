package tagModel

import (
	. "github.com/swsad-dalaotelephone/Server/database"
	"github.com/swsad-dalaotelephone/Server/modules/log"
)

// table name
const (
	TagTableName = "tag"
)

type Tag struct {
	Id   int    `gorm:"column:id; primary_key; not null" json:"tag_id"`
	Name string `gorm:"column:name; not null; unique"  json:"name"`
}

func initTag() {
	AddTag(Tag{Name: "运动"})
	AddTag(Tag{Name: "音乐"})
	AddTag(Tag{Name: "社会"})
	AddTag(Tag{Name: "效率"})
	AddTag(Tag{Name: "心理"})
	AddTag(Tag{Name: "美食"})
}

// if not exist table, create table
func init() {
	if !DB.HasTable(TagTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Tag{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
		initTag()
	}
}

//set table name
func (u Tag) TableName() string {
	return TagTableName
}

/*
 add new tag
 @parm new tag
 @return isSuccessful
*/
func AddTag(tag Tag) bool {
	DB.Create(&tag)
	res := DB.NewRecord(&tag) //return `false` after `tag` created
	return !res
}

// query tags by string key
func GetTagsByStrKey(key string, value string) (tags []Tag, err error) {
	res := DB.Where(key+" = ?", value).Find(&tags)
	err = res.Error
	return tags, err
}

// query tags by int key
func GetTagsByIntKey(key string, value int) (tags []Tag, err error) {
	res := DB.Where(key+" = ?", value).Find(&tags)
	err = res.Error
	return tags, err
}

/*
 update tag info
 must GetTagByKey first
*/
func UpdateTag(tag Tag) error {
	err := DB.Save(&tag).Error
	return err
}

/*
 get all tags
*/
func GetAllTags() (tags []Tag, err error) {
	err = DB.Find(&tags).Error
	return tags, err
}

/*
delete tag by id
*/
func DeleteTagById(id int) error {
	err := DB.Where("id = ?", id).Delete(Tag{}).Error
	return err
}

/*
delete tag by name
*/
func DeleteTagByName(name string) error {
	err := DB.Where("name = ?", name).Delete(Tag{}).Error
	return err
}

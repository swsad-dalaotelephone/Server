package tagModel

import (
	. "baobaozhuan/database"
)

// table name
const (
	TagTableName = "tag"
)

type Tag struct {
	Id   int    `gorm:"primary_key; not null" json:"tag_id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(TagTableName) {
		DB.CreateTable(&Tag{})
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

// query tag by key
func GetTagByKey(key string, value string) (Tag, error) {
	tag := Tag{}
	res := DB.Where(key+" = ?", value).First(&tag)
	err := res.Error
	return tag, err
}

/*
 update tag info
 must GetTagByKey first
*/
func UpdateTag(tag Tag) {
	DB.Save(&tag)
}

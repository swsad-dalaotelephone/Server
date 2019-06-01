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

/*
add all school of sysu
*/
func initSchool() {
	AddSchool(School{Name: "中国语言文学系", Type: "文科"})
	AddSchool(School{Name: "历史学系", Type: "文科"})
	AddSchool(School{Name: "社会学与人类学学院", Type: "文科"})
	AddSchool(School{Name: "博雅学院", Type: "文科"})
	AddSchool(School{Name: "岭南学院", Type: "经管"})
	AddSchool(School{Name: "外国语学院", Type: "文科"})
	AddSchool(School{Name: "法学院", Type: "文科"})
	AddSchool(School{Name: "政治与公共事物管理学院", Type: "文科"})
	AddSchool(School{Name: "管理学院", Type: "经管"})
	AddSchool(School{Name: "马克思主义学院", Type: "文科"})
	AddSchool(School{Name: "心理学系", Type: "理工"})
	AddSchool(School{Name: "传播与设计学院", Type: "文科"})
	AddSchool(School{Name: "资讯管理学院", Type: "经管"})
	AddSchool(School{Name: "艺术学院", Type: "文科"})
	AddSchool(School{Name: "数学学院", Type: "理工"})
	AddSchool(School{Name: "物理学院", Type: "理工"})
	AddSchool(School{Name: "化学学院", Type: "理工"})
	AddSchool(School{Name: "地理科学与规划学院", Type: "理工"})
	AddSchool(School{Name: "生命科学学院", Type: "理工"})
	AddSchool(School{Name: "工学院", Type: "理工"})
	AddSchool(School{Name: "材料科学与工程学院", Type: "理工"})
	AddSchool(School{Name: "电子与信息工程学院", Type: "理工"})
	AddSchool(School{Name: "数据科学与计算机学院", Type: "理工"})
	AddSchool(School{Name: "环境科学与工程学院", Type: "理工"})
	AddSchool(School{Name: "中山医学院", Type: "医学"})
	AddSchool(School{Name: "光华口腔医学院", Type: "医学"})
	AddSchool(School{Name: "公共卫生学院", Type: "医学"})
	AddSchool(School{Name: "药学院", Type: "医学"})
	AddSchool(School{Name: "护理学院", Type: "医学"})
	AddSchool(School{Name: "中国语言文学系（珠海）", Type: "文科"})
	AddSchool(School{Name: "历史学系（珠海）", Type: "文科"})
	AddSchool(School{Name: "哲学系（珠海）", Type: "文科"})
	AddSchool(School{Name: "国际金融学院", Type: "经管"})
	AddSchool(School{Name: "国际翻译学院", Type: "文科"})
	AddSchool(School{Name: "国际关系学院", Type: "文科"})
	AddSchool(School{Name: "旅游学院", Type: "经管"})
	AddSchool(School{Name: "数学学院（珠海）", Type: "理工"})
	AddSchool(School{Name: "物理与天文学院", Type: "理工"})
	AddSchool(School{Name: "大气科学学院", Type: "理工"})
	AddSchool(School{Name: "海洋科学学院", Type: "理工"})
	AddSchool(School{Name: "地理科学与工程学院", Type: "理工"})
	AddSchool(School{Name: "化学工程与技术学院", Type: "理工"})
	AddSchool(School{Name: "海洋工程与技术学院", Type: "理工"})
	AddSchool(School{Name: "中法核工程与技术学院", Type: "理工"})
	AddSchool(School{Name: "土木工程学院", Type: "理工"})
	AddSchool(School{Name: "医学院", Type: "医学"})
	AddSchool(School{Name: "公共卫生学院（深圳）", Type: "医学"})
	AddSchool(School{Name: "药学院（深圳）", Type: "医学"})
	AddSchool(School{Name: "材料学院", Type: "理工"})
	AddSchool(School{Name: "生物医学工程学院", Type: "理工"})
	AddSchool(School{Name: "电子与通信工程学院", Type: "理工"})
	AddSchool(School{Name: "智能工程学院", Type: "理工"})
	AddSchool(School{Name: "航空航天学院", Type: "理工"})
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
func AddSchool(school School) bool {
	DB.Create(&school)
	res := DB.NewRecord(&school) //return `false` after `school` created
	return !res
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

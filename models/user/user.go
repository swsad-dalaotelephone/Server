package userModel

import (
	"time"

	"github.com/swsad-dalaotelephone/Server/models/preference"

	. "github.com/swsad-dalaotelephone/Server/database"

	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/jinzhu/gorm"
)

// table name
const (
	UserTableName = "user"
)

type User struct {
	Id          string                       `gorm:"column:id; type:varchar(36); primary_key; not null" json:"id"`
	OpenId      string                       `gorm:"column:open_id; type:varchar(64); index: open_id_idx" json:"open_id"`
	NickName    string                       `gorm:"column:nick_name; type:varchar(64)" json:"nick_name"`
	Password    string                       `gorm:"column:password; type:varchar(64)" json:"password"`
	Phone       string                       `gorm:"column:phone; type:varchar(20); unique_index: phone_idx; not null" json:"phone"`
	Birthday    time.Time                    `gorm:"column:birtyday" json:"birthday"`
	Gender      int                          `gorm:"column:gender" json:"gender"` //male - 1   female  - 2  empty - 0
	CampusId    int                          `gorm:"column:campus_id" json:"campus_id"`
	SchoolId    string                       `gorm:"column:school_id" json:"school_id"`
	Grade       string                       `gorm:"column:grade; type:varchar(20);" json:"grade"`
	Account     int                          `gorm:"column:account; default:0" json:"account"`
	Email       string                       `gorm:"column:email" json:"email"`
	Evaluation  int                          `gorm:"column:evaluation" json:"evaluation"`
	CreatedAt   time.Time                    `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time                    `gorm:"column:updated_at" json:"-"`
	Preferences []preferenceModel.Preference `gorm:"foreignkey:UserId"`
}

// if not exist table, create table
func init() {
	if !DB.HasTable(UserTableName) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
			log.ErrorLog.Println(err)
		}
	}
}

//set table name
func (u User) TableName() string {
	return UserTableName
}

// set user.id as uuid before create
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.Id = util.GetNewUuid()
	user.Birthday = time.Now()
	return nil
}

/*
 add new user
 @parm new user
 @return isSuccessful
*/
func AddUser(user User) (User, bool) {
	DB.Create(&user)
	res := DB.NewRecord(&user) //return `false` after `user` created
	return user, !res
}

// query users by string key
func GetUsersByStrKey(key string, value string) (users []User, err error) {
	err = DB.Where(key+" = ?", value).Find(&users).Error
	return users, err
}

// query users by int key
func GetUsersByIntKey(key string, value int) (users []User, err error) {
	err = DB.Where(key+" = ?", value).Find(&users).Error
	return users, err
}

/*
 update user info
 must GetUserByKey first
*/
func UpdateUser(user User) error {
	err := DB.Save(&user).Error
	return err
}

/*
 get all users
*/
func GetAllUsers() (users []User, err error) {
	err = DB.Find(&users).Error
	return users, err
}

/*
delete user by id
*/
func DeleteUserById(id string) error {
	err := DB.Where("id = ?", id).Delete(User{}).Error
	return err
}

/*
get user's preferences
*/
func GetPreferenceByUser(user User) (preferences []preferenceModel.Preference, err error) {
	err = DB.Model(&user).Related(&preferences).Error
	return preferences, err
}

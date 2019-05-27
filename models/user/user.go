package userModel

import (
	. "github.com/swsad-dalaotelephone/Server/database"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

const (
	UserTableName = "user"
)

type User struct {
	Id        string     `gorm:"primary_key" json:"id"`
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Gender    int        `json:"gender"` //male - 1   female  - 2  empty - 0
	Email     string     `json:"email"`
	Birthday  time.Time  `json:"birthday"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

// func init() {
// 	if !DB.HasTable(UserTableName) {
// 		DB.Table(UserTableName).CreateTable(&User{})
// 	}
// }

// set user.id as uuid before create
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.Id = uuid.Must(uuid.NewV4(), nil).String()
	return nil
}

/*
 add new user with (openid,nickName,Password,phone)
 @parm infomation of user
 @return user, isSuccessful
*/
func AddUser(openId, nickName, password, phone string) (User, bool) {
	// set current time to birthday
	user := User{OpenId: openId, NickName: nickName, Password: password, Phone: phone, Birthday: time.Now()}
	DB.Table(UserTableName).Create(&user)
	res := DB.Table(UserTableName).NewRecord(&user) //return `false` after `user` created
	return user, !res
}

// query user by openid
func GetUserByKey(key string, value string) (User, error) {
	user := User{}
	res := DB.Table(UserTableName).Where(key+" = ?", value).First(&user)
	err := res.Error
	return user, err
}

// update user info
func UpdateUser(id, nickName, password, phone string) User {
	user := User{}
	DB.Table(UserTableName).Where("Id = ?", id).First(&user)
	user.NickName = nickName
	user.Password = password
	user.Phone = phone

	DB.Table(UserTableName).Save(&user)
	return user
}

package userModel

import "testing"
import . "github.com/swsad-dalaotelephone/Server/database"

func TestAddUser(t *testing.T) {
	AddUser(User{NickName: "xxx", Password: "ttt", Phone: "12312312311"})
}

func TestGetUser(t *testing.T) {
	users1, err1 := GetUsersByStrKey("phone", "13")
	t.Log(users1, err1)
	users2, err2 := GetUsersByStrKey("phone", "12312312311")
	t.Log(users2, err2)
}

func TestUpdateUser(t *testing.T) {
	users1, err1 := GetUsersByStrKey("phone", "12312312311")
	t.Log(err1)
	user := users1[0]
	user.Password = "yyy"
	err := UpdateUser(user)
	t.Log(err)
	users1, err1 = GetUsersByStrKey("phone", "12312312311")
	t.Log(users1, err1)
}

func TestRelateQuery(t *testing.T) {
	users, _ := GetUsersByStrKey("phone", "12312312311")
	AddPreference(Preference{UserId: users[0].Id, TagId: 11})
	preferences, _ := GetPreferenceByUser(users[0])
	t.Log(preferences)
	DB.Model(&users[0]).Association("Preferences").Append(Preference{TagId: 12})
	DB.Model(&users[0]).Association("Preferences").Find(&preferences)
	t.Log(preferences)
	t.Log(users[0])
}

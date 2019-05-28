package userModel

import "testing"

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

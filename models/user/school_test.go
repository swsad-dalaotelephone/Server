package userModel

import "testing"

func TestAddSchool(t *testing.T) {
	school1 := School{}
	school2 := School{Name: "数据科学与计算机学院", Type: "理工"}
	school3 := School{Name: "管理学院", Type: "商科"}
	res1 := AddSchool(school1)
	res2 := AddSchool(school2)
	res3 := AddSchool(school3)
	t.Log(res1)
	t.Log(res2)
	t.Log(res3)
}
func TestGetSchool(t *testing.T) {
	schools1, err1 := GetSchoolsByStrKey("type", "理工")
	schools2, err2 := GetSchoolsByStrKey("type", "商科")
	t.Log(err1)
	t.Log(schools1[0])
	t.Log(err2)
	t.Log(schools2[0])
}

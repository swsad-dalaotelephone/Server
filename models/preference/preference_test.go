package preferenceModel

import "testing"

func TestAddPreference(t *testing.T) {
	preference, res := AddPreference(Preference{UserId: "abc", TagId: 11})
	t.Log(preference, res)
	preference, res = AddPreference(Preference{UserId: "abc", TagId: 2})
	t.Log(preference, res)
	preference, res = AddPreference(Preference{UserId: "aaa", TagId: 11})
	t.Log(preference, res)
	preference, res = AddPreference(Preference{UserId: "aaa", TagId: 2})
	t.Log(preference, res)
}

func TestGetPreference(t *testing.T) {
	preference1, _ := GetPreferencesByStrKey("user_id", "abc")
	t.Log(preference1)
	preference2, _ := GetPreferencesByIntKey("tag_id", 11)
	t.Log(preference2)
}

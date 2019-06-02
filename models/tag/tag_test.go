package tagModel

import "testing"

func TestAddTag(t *testing.T) {
	tag, res := AddTag(Tag{Name: "美食"})
	t.Log(tag, res)
}
func TestTag(t *testing.T) {
	tags, _ := GetAllTags()
	t.Log(tags)
}

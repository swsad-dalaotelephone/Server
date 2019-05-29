package tagModel

import "testing"

func TestAddTag(t *testing.T) {

}
func TestTag(t *testing.T) {
	tags, _ := GetAllTags()
	t.Log(tags)
}

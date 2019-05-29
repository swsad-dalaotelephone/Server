package taskModel

import (
	"testing"
)

func TestAddAcceptance(t *testing.T) {
	AddAcceptance(Acceptance{TaskId: "11", AccepterId: "xxx"})
	acc, err := GetAcceptanceByTaskAccepterId("1", "x")
	t.Log(acc, err)
	acc, err = GetAcceptanceByTaskAccepterId("11", "xxx")
	t.Log(acc, err)
}

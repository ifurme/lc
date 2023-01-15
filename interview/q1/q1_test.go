package q1

import "testing"

func TestIsUnique(t *testing.T) {
	s := "leetcode"
	b := isUnique(s)
	if b {
		t.Errorf("wrong")
	}
}

package avltree

import "testing"

func Compare(a, b interface{}) int {
	_a := a.(int)
	_b := b.(int)
	if int(_a) > int(_b) {
		return 1
	} else if int(_a) == int(_b) {
		return 0
	}
	return -1
}

func TestInsert(t *testing.T) {
	tree := New(Compare)
	tree.Insert(43)
}

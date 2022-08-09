package simplemath

import "testing"

// 对add的单元测试，go的命名规范
func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, excepted 3.", r)
	}
}

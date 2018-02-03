package hash

import "testing"

func TestChHash(t *testing.T) {
	hash := NewChHash(16)
	hash.Set(StringHashFunc("One"), 2)
	val, ok := hash.Get(StringHashFunc("One"))
	if !ok {
		t.Error("Hash表异常，无法获取已存在的键")
	}
	if val != 2 {
		t.Error("Hash表异常，获取跟设置不一样")
	}

	hash.Set(StringHashFunc("One"), 5)
	val, _ = hash.Get(StringHashFunc("One"))
	if val != 5 {
		t.Error("Hash表异常， 无法重新设置值")
	}

	hash.Delete(StringHashFunc("One"))
	val, ok = hash.Get(StringHashFunc("One"))
	if ok {
		t.Error("Hash表异常， 删除失败")
	}
}

package hash

import "testing"

var tests = []struct {
	key string
	val int
}{
	{"one", 2},
	{"two", 4},
	{"hello", 7},
	{"uuid", 0},
	{"3434", 122},
}

func TestChHash(t *testing.T) {
	hash := NewChHash(5)
	testInsertDelete(hash, t)
}

func TestOHash(t *testing.T) {
	hash := NewOHash(5)
	testInsertDelete(hash, t)
}

func TestOHashExceedInsert(t *testing.T) {
	h := NewOHash(5)
	for _, val := range tests {
		h.Set(StringHashFunc(val.key), val.val)
	}
	err := h.Set(StringHashFunc("six"), 12)
	if err != ErrBucketFull {
		t.Error("OHash 插入异常， 队列已满, 不能再插入")
	}
}

func testInsertDelete(hash Tabler, t *testing.T) {
	for _, data := range tests {
		func() {
			k := StringHashFunc(data.key)
			// test set
			hash.Set(k, data.val)
			val, ok := hash.Get(k)
			if !ok {
				t.Error("Hash表异常，无法获取已存在的键")
			}
			if val != data.val || hash.Len() != 1 {
				t.Error("Hash表异常，获取跟设置不一样")
			}

			// test reset
			hash.Set(k, 5)
			val, _ = hash.Get(k)
			if val != 5 || hash.Len() != 1 {
				t.Error("Hash表异常， 无法重新设置值")
			}

			// test delete
			hash.Delete(k)
			val, ok = hash.Get(k)
			if ok || hash.Len() != 0 {
				t.Error("Hash表异常， 删除失败")
			}
		}()
	}
}

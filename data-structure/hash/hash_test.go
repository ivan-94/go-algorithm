package hash

import (
	"fmt"
	"testing"
)

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

func TestGrowHashInsert(t *testing.T) {
	hash := NewGrowHash()
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

func TestGrowHash(t *testing.T) {
	keys := []string{}
	for i := 0; i < 60; i++ {
		keys = append(keys, fmt.Sprint(i))
	}
	h := NewGrowHash()
	// set
	for i, key := range keys {
		h.Set(StringHashFunc(key), key)
		if h.Len() != i+1 {
			t.Errorf("GrowHash 插入失败，长度没有变化, 期望%d, 实际上是 %d\n", i+1, h.Len())
		}
	}

	if h.count != 16 {
		t.Errorf("GrowHash 自动增长失败, 期望容量为%d, 实际上是 %d\n", 16, h.count)
	}

	// reset
	for _, key := range keys {
		val := key + "++"
		h.Set(StringHashFunc(key), val)
		data, ok := h.Get(StringHashFunc(key))
		if !ok || data != val {
			t.Errorf("GrowHash 赋值失败, 期望结果%s, 实际上是 %s\n", val, data)
		}
	}

	if h.count != 16 || h.oldbuckets != nil {
		t.Errorf("排空失败")
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

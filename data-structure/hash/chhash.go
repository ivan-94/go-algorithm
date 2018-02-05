// Package hash 实现hash数据结构相关算法
package hash

import (
	"github.com/carney520/go-algorithm/data-structure/list"
)

// Match 用于比较两个键是否相等
type Match func(a, b interface{}) bool

// Hasher 表示一个可求hash的接口
// 所有键必须实现这个接口
type Hasher interface {
	hash() int
}

// Tabler 表示不同类型hash表需要实现的方法
type Tabler interface {
	Get(key Hasher) (val interface{}, ok bool)
	Set(key Hasher, val interface{}) error
	Len() int
	Delete(key Hasher)
}

// ChHash 表示链式hash表
type ChHash struct {
	count int          // 桶数量
	table []*list.List // 桶列表
	len   int          // 数据长度
}

type item struct {
	key Hasher
	val interface{}
}

// 查找
func (h *ChHash) lookup(key Hasher) (entry *item, node *list.Node, bucket *list.List) {
	// 找到桶的索引
	i := key.hash() % h.count
	bucket = h.table[i]
	var data *item
	bucket.Each(func(n *list.Node, index int) bool {
		val, ok := n.Data.(*item)
		if ok {
			if val.key == key {
				data = val
				node = n
				return true
			}
		}
		return false
	})

	return data, node, bucket
}

// Get 从hash中获取数据
func (h *ChHash) Get(key Hasher) (val interface{}, ok bool) {
	i, _, _ := h.lookup(key)
	if i != nil {
		return i.val, true
	}
	return nil, false
}

// Len 获取hash的数据项数量
func (h *ChHash) Len() int {
	return h.len
}

// Delete 删除指定键
func (h *ChHash) Delete(key Hasher) {
	i, node, bucket := h.lookup(key)
	if i != nil {
		bucket.Remove(node)
		h.len--
	}
}

// Set 存储一个值
func (h *ChHash) Set(key Hasher, val interface{}) (err error) {
	i, _, _ := h.lookup(key)
	if i != nil {
		i.val = val
		return
	}
	id := key.hash() % h.count
	bucket := h.table[id]
	i = new(item)
	i.key = key
	i.val = val
	bucket.Append(i)
	h.len++
	return
}

// NewChHash 创建一个链式哈希表
func NewChHash(size int) *ChHash {
	h := &ChHash{
		count: size,
		table: make([]*list.List, size),
	}
	for i := 0; i < size; i++ {
		h.table[i] = list.New()
	}
	return h
}

// StringHash 实现Hasher接口， 用于计算字符串的hash
type StringHash string

func (s StringHash) hash() int {
	var val int
	for i := 0; i < len(s); i++ {
		val = (val << 4) + int(s[i])
		if tmp := val & 0xf0000000; tmp != 0 {
			val = val ^ (tmp >> 24)
			val = val ^ tmp
		}
	}
	return val
}

// StringHashFunc 将string转换成StringHash
func StringHashFunc(str string) StringHash {
	return StringHash(str)
}

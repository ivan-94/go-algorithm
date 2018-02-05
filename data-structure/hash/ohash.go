package hash

import "errors"

// 错误代码
var (
	ErrBucketFull = errors.New("Buckets exceed")
	ErrUnknown    = errors.New("Unknown error")
)

// OHash 表示开地址hash表
type OHash struct {
	count int     // 桶数量
	table []*item // 桶列表
	len   int     // 数据量
}

func (h *OHash) lookup(key Hasher) (data *item, pos int) {
	k := key.hash()
	pos = -1
	for i := 0; i < h.count; i++ {
		pos = (h.h1(k) + i*h.h2(k)) % h.count
		if h.table[pos] == nil {
			return
		} else if h.table[pos].key == key {
			return h.table[pos], pos
		}
	}
	return
}

func (h *OHash) h1(k int) int {
	return k % h.count
}

func (h *OHash) h2(k int) int {
	return k % (h.count - 2)
}

// Len 返回hash的数量
func (h *OHash) Len() int {
	return h.len
}

// Set 存储一个值
func (h *OHash) Set(key Hasher, val interface{}) error {
	// 不能再插入了
	if h.len == h.count {
		return ErrBucketFull
	}
	it, pos := h.lookup(key)
	if it != nil {
		// reset
		it.val = val
	} else if pos != -1 {
		// new
		h.table[pos] = &item{
			key: key,
			val: val,
		}
		h.len++
	} else {
		return ErrUnknown
	}
	return nil
}

// Delete 删除指定键
func (h *OHash) Delete(key Hasher) {
	it, pos := h.lookup(key)
	if it != nil && pos != -1 {
		h.table[pos] = nil
		h.len--
	}
}

// Get 获取指定键的值
func (h *OHash) Get(key Hasher) (val interface{}, ok bool) {
	it, _ := h.lookup(key)
	if it != nil {
		return it.val, true
	}
	return nil, false
}

// NewOHash 创建一个hash表
// size 要求是一个素数
func NewOHash(size int) *OHash {
	return &OHash{
		count: size,
		table: make([]*item, size),
	}
}

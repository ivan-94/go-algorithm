package hash

// TODO:
// 这个文件实现了自动增长的Hash表
// go 原生的map类型就是自增长的：
// > go map类型每个bucket包含8个键值对。
// > 哈希值的低端位用于确定选择哪个bucket。
// > 每个bucket包含了一些hash值得高端位，用于区分每个数据项在
// > bucket中的存储;
// >
// > 如果bucket超过8个键值对, go会创建一个新的列表连接到这个bucket
// >
// > 当hash表增长时， go会分配一个新的buckets数组， 新的buckets是旧的buckets的
// > 两倍。 并使用增量的形式将旧的buckets数据拷贝到新的buckets中
import (
	"github.com/carney520/go-algorithm/data-structure/list"
)

const (
	loadFactorNum = 13
	loadFactorDen = 2
)

// GrowHash 自增长的hash
type GrowHash struct {
	b          uint
	count      int          // 大小
	buckets    []*list.List // 桶列表
	oldbuckets []*list.List // 旧桶列表
	len        int          // 数据量
}

func find(bucket *list.List, key Hasher) (data *item, node *list.Node) {
	bucket.Each(func(n *list.Node, index int) bool {
		val, ok := n.Data.(*item)
		if ok && val.key == key {
			// find it
			data = val
			node = n
			return true
		}
		return false
	})
	return
}

func (h *GrowHash) lookup(key Hasher) (data *item, node *list.Node) {
	i := key.hash()
	// 从旧bucket中检索
	if h.oldbuckets != nil {
		hash := i % len(h.oldbuckets)
		bucket := h.oldbuckets[hash]
		if bucket != nil {
			// find here
			data, node = find(bucket, key)
			if data != nil {
				return
			}
		}
	}

	// 从新bucket中检索
	hash := i % len(h.buckets)
	bucket := h.buckets[hash]
	if bucket != nil {
		data, node = find(bucket, key)
	}
	return
}

// Get 获取指定键的值
func (h *GrowHash) Get(key Hasher) (val interface{}, ok bool) {
	it, _ := h.lookup(key)
	if it != nil {
		return it.val, true
	}
	return nil, false
}

// Len 获取hash的数据长度
func (h *GrowHash) Len() int {
	return h.len
}

// Set 设置值
func (h *GrowHash) Set(key Hasher, val interface{}) (err error) {
	i := key.hash()

again:
	// 判断是否存在旧列表
	if h.oldbuckets != nil {
		hash := i % len(h.oldbuckets)
		bucket := h.oldbuckets[hash]
		if bucket != nil {
			// 将旧的bucket迁移到新的桶列表
			h.reHash(hash)
		}
	}

	if h.buckets == nil {
		h.buckets = make([]*list.List, h.count)
	}

	// 插入
	hash := i % len(h.buckets)
	bucket := h.buckets[hash]
	if bucket == nil {
		bucket = list.New()
		h.buckets[hash] = bucket
	}
	if bucket.Len() > 0 {
		data, _ := find(bucket, key)
		if data != nil {
			data.val = val
			return
		}
	}

	bucket.Append(&item{
		key: key,
		val: val,
	})
	h.len++

	// 判断是否要扩容
	if h.oldbuckets == nil && h.overLoadFactory() {
		h.hashGrow()
		goto again
	}

	return
}

// Delete 删除指定键
func (h *GrowHash) Delete(key Hasher) {
	i := key.hash()

	if h.oldbuckets != nil {
		hash := i % len(h.oldbuckets)
		bucket := h.oldbuckets[hash]
		if bucket != nil {
			// 将旧的bucket迁移到新的桶列表
			h.reHash(hash)
		}
	}

	if h.buckets == nil {
		return
	}

	// 删除
	hash := i % len(h.buckets)
	bucket := h.buckets[hash]
	if bucket != nil {
		item, node := find(bucket, key)
		if item != nil {
			// find it
			bucket.Remove(node)
			h.len--
		}
	}
}

func (h *GrowHash) hashGrow() {
	h.oldbuckets = h.buckets
	h.b++
	h.count = 1 << h.b
	h.buckets = make([]*list.List, h.count)
}

// 迁移到新的桶列表
func (h *GrowHash) reHash(index int) {
	bucket := h.oldbuckets[index]
	bucket.Each(func(n *list.Node, _ int) bool {
		val, ok := n.Data.(*item)
		if ok {
			hash := val.key.hash() % len(h.buckets)
			newbucket := h.buckets[hash]
			if newbucket == nil {
				newbucket = list.New()
				h.buckets[hash] = newbucket
			}
			newbucket.Append(n.Data)
		}
		return false
	})
	h.oldbuckets[index] = nil
	if h.evacuated() {
		h.oldbuckets = nil
	}
}

func (h *GrowHash) evacuated() bool {
	for _, bucket := range h.oldbuckets {
		if bucket != nil {
			return false
		}
	}
	return true
}

func (h *GrowHash) overLoadFactory() bool {
	l := loadFactorNum * (h.count / loadFactorDen)
	return h.len > l
}

// NewGrowHash 创建一个自增的hash表
func NewGrowHash() *GrowHash {
	return &GrowHash{
		b:     3,
		count: 1 << 3,
	}
}

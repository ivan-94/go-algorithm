package list

// 双向列表

// DNode 是双向链表节点
type DNode struct {
	Data interface{}
	prev *DNode
	next *DNode
	list *DList
}

// DList 代表双向列表
type DList struct {
	head *DNode
	tail *DNode
	len  int
}

// Len 获取链表长度
func (l *DList) Len() int {
	return l.len
}

// Append 追加一个元素到列表中， 如果提供的前继， 则追加到前继知乎
func (l *DList) Append(data interface{}, n ...*DNode) (*DNode, error) {
	nn := &DNode{Data: data, list: l}
	if len(n) > 0 {
		el := n[0]
		if el.list != l {
			return nil, ErrNotFound
		}
		nn.next = el.next
		nn.prev = el
		el.next.prev = nn
		el.next = nn
		if el == l.tail {
			l.tail = nn
		}
	} else if l.head == nil {
		// empty list
		l.head = nn
		l.tail = nn
	} else {
		// append to tail
		l.tail.next = nn
		nn.prev = l.tail
		l.tail = nn
	}
	l.len++
	return nn, nil
}

// Prepend 在头部或者在指定节点之前插入
func (l *DList) Prepend(data interface{}, n ...*DNode) (*DNode, error) {
	nn := &DNode{Data: data, list: l}
	if len(n) > 0 {
		// 插入到el之前
		el := n[0]
		if el.list != l {
			return nil, ErrNotFound
		}
		if prev := el.prev; prev != nil {
			nn.prev = prev
			nn.next = el
			el.prev.next = nn
			el.prev = nn
		} else {
			// 没有前继， 那就是head
			l.head = nn
			nn.next = el
		}
	} else if l.head == nil {
		l.head = nn
		l.tail = nn
	} else {
		// 插入到head之前
		nn.next = l.head
		l.head.prev = nn
		l.head = nn
	}
	l.len++

	return nn, nil
}

// Remove 从链表中移除节点
func (l *DList) Remove(n *DNode) error {
	if n.list == nil {
		// removed
		return nil
	}

	if n.list != l {
		return ErrNotFound
	}

	if prev := n.prev; prev != nil {
		// 中间节点或结尾节点
		prev.next = n.next
		if n.next != nil {
			n.next.prev = prev
		} else {
			l.tail = prev
		}
	} else if l.head == n {
		// 头部
		if n.next != nil {
			l.head = n.next
		} else {
			// 最后一个元素
			l.head = nil
			l.tail = nil
		}
	}

	l.len--
	// reset node
	n.list = nil
	n.prev = nil
	n.next = nil
	return nil
}

// Each 顺序迭代
func (l *DList) Each(it func(data interface{}, index int) bool) {
	if l.len == 0 {
		return
	}
	cur, i := l.head, 0
	for cur != nil {
		stop := it(cur.Data, i)
		if stop {
			break
		}
		cur = cur.next
		i++
	}
}

// RvEach 反序迭代
func (l *DList) RvEach(it func(data interface{}, index int) bool) {
	if l.len == 0 {
		return
	}
	cur, i := l.tail, 0
	for cur != nil {
		stop := it(cur.Data, i)
		if stop {
			break
		}
		cur = cur.prev
		i++
	}
}

// Has 确认节点是否在链表内
func (l *DList) Has(n *DNode) bool {
	return n.list == l
}

// NewDList 创建一个双向链表
func NewDList() *DList {
	return &DList{}
}

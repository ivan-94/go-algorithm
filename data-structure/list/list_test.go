package list

// 单链表测试

import "testing"

func TestListAppend(t *testing.T) {
	l := New()
	h, _ := l.Append("hello")

	if l.Len != 1 {
		t.Errorf("链表长度有误, 期望%d, 实际上是%d", 1, l.Len)
	}

	if l.head.Data != "hello" {
		t.Errorf("链表数据存储有误, 实际上为: %v", l.head.Data)
	}

	if l.head != h || l.tail != h {
		t.Errorf("链表插入异常, 插入空链表, 首位应该是同一个元素")
	}

	n, _ := l.Append("world")

	if l.Len != 2 {
		t.Errorf("链表长度有误, 期望%d, 实际上是%d", 2, l.Len)
	}

	if l.head.next.Data != "world" {
		t.Errorf("链表数据存储有误, 实际上为: %v", l.head.next.Data)
	}

	nn, _ := n.Append("foo")
	if l.Len != 3 {
		t.Errorf("链表长度有误, 期望%d, 实际上是%d", 3, l.Len)
	}

	if l.tail != nn {
		t.Errorf("链表插入异常, 链表尾部不等于最新元素")
	}

	if n.next != nn {
		t.Errorf("链表插入异常, 元素下一个元素不等于最新元素")
	}

	if !l.Has(nn) {
		t.Errorf("nn 应该是l的节点")
	}

	notExisted := Node{}
	_, err := notExisted.Append("a")
	if err != ErrRemoved {
		t.Errorf("链表插入异常, 悬空节点无法被追加元素")
	}
}

func TestListEach(t *testing.T) {
	l := New()
	l.Append(1)
	n, _ := l.Append(2)
	l.Append(3)
	n.Append(4)
	var act [4]int
	exp := [...]int{1, 2, 4, 3}
	l.Each(func(d interface{}, i int) bool {
		v := d.(int)
		act[i] = v
		return false
	})

	if act != exp {
		t.Errorf("遍历异常, 期望")
	}
}

func TestListRemove(t *testing.T) {
	l := New()
	foo, _ := l.Append("1")
	bar, _ := l.Append("2")
	baz, _ := l.Append("3")

	bar.Remove()
	if bar.list != nil {
		t.Errorf("删除后应该为悬空节点")
	}

	if foo.next != baz {
		t.Errorf("foo的后继应该是baz")
	}

	foo.Remove()
	if l.head != baz || l.tail != baz {
		t.Errorf("foo删除后, baz是最后一个元素, 应该是head和tail")
	}

	bazz, _ := baz.Append("4")
	if l.tail != bazz {
		t.Errorf("bazz应该是最后一个元素")
	}

	bazz.Remove()

	if l.head != baz || l.tail != baz {
		t.Errorf("bazz删除后, baz是最后一个元素, 应该是head和tail")
	}

	baz.Remove()
	if l.head != nil || l.tail != nil {
		t.Errorf("baz删除后, head, tail 应该为空")
	}
}

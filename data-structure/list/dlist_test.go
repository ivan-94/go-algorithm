package list

import "testing"

func TestDListAppend(t *testing.T) {
	l := NewDList()
	n, _ := l.Append("1")
	if l.Len() != 1 {
		t.Error("插入异常, 链表长度有误")
	}
	if l.head != n || l.tail != n {
		t.Error("插入异常, 空链表head和tail不为新插入元素")
	}

	nn, _ := l.Append("2")
	if nn.prev != n {
		t.Error("插入异常, 前继异常")
	}

	if n.next != nn {
		t.Error("插入异常, 后继异常")
	}

	if l.tail != nn {
		t.Error("插入异常, tail 部位新插入元素")
	}

	nnn, _ := l.Append("3", n)
	if nnn.next != nn {
		t.Error("插入异常, nnn的后继应该为n的前继")
	}

	if nn.prev != nnn {
		t.Error("插入异常, nn的前继应该为nnn")
	}

	if n.next != nnn {
		t.Error("插入异常, 插入元素应该为被插入元素的后继")
	}

	if nnn.prev != n {
		t.Error("插入异常, 被插入元素应该为插入元素的前继")
	}
}

func TestDListPrepend(t *testing.T) {
	l := NewDList()
	n, _ := l.Prepend("1")
	if l.Len() != 1 || l.head != n || l.tail != n {
		t.Errorf("插入异常，插入第一个元素应该等于head和tail")
	}

	nn, _ := l.Prepend("2")
	if l.Len() != 2 || l.head != nn || l.tail != n || n.prev != nn || nn.next != n {
		t.Errorf("插入异常，插入第二个元素应该等于head， 并且在第一个元素之前")
	}

	nnn, _ := l.Prepend("3", n)
	if l.Len() != 3 || nn.next != nnn || nnn.prev != nn || nnn.next != n || n.prev != nnn {
		t.Errorf("插入异常，插入第三个元素在最后一个元素之前")
	}
}

func TestDListRemove(t *testing.T) {
	l := NewDList()
	foo, _ := l.Append("1")
	bar, _ := l.Append("2")
	baz, _ := l.Append("3")

	l.Remove(bar)
	if foo.next != baz {
		t.Errorf("删除异常， foo的后继应该为baz")
	}
	if baz.prev != foo {
		t.Errorf("删除异常， baz的前继应该为foo")
	}

	if l.Has(bar) {
		t.Errorf("删除异常， bar 应该不在链表内")
	}

	if l.Len() != 2 {
		t.Errorf("删除异常， 删除后长度没有递减")
	}

	l.Remove(baz)
	if l.head != foo || l.tail != foo {
		t.Errorf("删除异常， foo是最后一个元素， 应该是head和tail")
	}

	l.Remove(foo)
	if l.head != nil || l.tail != nil {
		t.Errorf("删除异常， 链表为空时，head 和tail应该为nil")
	}

	if l.Len() != 0 {
		t.Errorf("链表为空时， len不为0")
	}
}

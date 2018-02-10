package bintree

import "testing"

func TestBinTreeInsertLeft(t *testing.T) {
	tree := BinTree{}
	n, err := tree.InsLeft("hello")
	if n.Data != "hello" || err != nil {
		t.Errorf("插入到左节点失败, 返回值错误")
	}
	if tree.Len() != 1 {
		t.Errorf("插入到左节点失败, 树长度错误")
	}

	tree.InsLeft("world", n)
	if tree.Len() != 2 || n.left.Data != "world" {
		t.Errorf("插入到左节点失败, 没有插入到指定节点的左节点")
	}
}

func TestBinTreeInsertRight(t *testing.T) {
	tree := BinTree{}
	n, err := tree.InsRight("hello")
	if n.Data != "hello" || err != nil {
		t.Errorf("插入到右节点失败, 返回值错误")
	}
	if tree.Len() != 1 {
		t.Errorf("插入到右节点失败, 树长度错误")
	}

	tree.InsRight("world", n)
	if tree.Len() != 2 || n.right.Data != "world" {
		t.Errorf("插入到右节点失败, 没有插入到指定节点的左节点")
	}
}

func TestBinTreeRemoveRoot(t *testing.T) {
	tree := BinTree{}
	tree.InsLeft("root")
	tree.InsLeft("left")
	tree.InsRight("right")
	node, err := tree.RmLeft()
	if err != nil || node.Data != "left" || tree.Len() != 2 {
		t.Errorf("移除根节点的左节点失败")
	}

	node, err = tree.RmRight()
	if err != nil || node.Data != "right" || tree.Len() != 1 {
		t.Errorf("移除根节点的右节点失败")
	}

	node, err = tree.RmLeft()
	if err != nil || node.Data != "root" || tree.Len() != 0 {
		t.Errorf("移除根节点失败")
	}

	node, err = tree.RmLeft()
	if err != ErrEmpty {
		t.Errorf("移除异常, 不能移除空节点")
	}
}

func TestBinTreeRemove(t *testing.T) {
	tree := BinTree{}
	r, _ := tree.InsLeft("root")
	l1, _ := tree.InsLeft("1", r)
	r1, _ := tree.InsRight("2", r)
	l1l, _ := tree.InsLeft("3", l1)
	l1r, _ := tree.InsRight("4", l1)

	if r.left != l1 || l1.left != l1l || l1.right != l1r || r.right != r1 {
		t.Error("插入失败, 插入位置错误")
	}

	if r.IsLeaf() || l1.IsLeaf() || !l1l.IsLeaf() || !l1r.IsLeaf() || !r1.IsLeaf() {
		t.Error("叶子节点判断有误")
	}

	if tree.len != 5 {
		t.Error("插入失败， 长度应该为5")
	}
	rm, _ := tree.RmLeft(r)
	if tree.len != 2 || rm != l1 {
		t.Errorf("删除失败， 递归删除子节点，长度应该为3, 实际上为: %d", tree.len)
	}
	tree.RmRight(r)
	if tree.len != 1 {
		t.Error("删除右节点失败")
	}
	tree.RmLeft()
	if tree.len != 0 {
		t.Error("删除根节点失败")
	}
}

func TestBinTreeClear(t *testing.T) {
	tree := BinTree{}
	check := func(index int) {
		if tree.root != nil || tree.len != 0 {
			t.Fatalf("test %d: 清除失败, tree.len为 %d", index, tree.len)
		}
	}
	tree.InsLeft("root")
	tree.Clear()
	check(1)

	tree = BinTree{}
	tree.InsLeft("root")
	tree.InsLeft("left")
	tree.Clear()
	check(2)

	tree = BinTree{}
	tree.InsLeft("root")
	tree.InsRight("right")
	tree.Clear()
	check(3)

	tree = BinTree{}
	tree.InsLeft("root")
	tree.InsRight("right")
	tree.InsLeft("left")
	tree.Clear()
	check(4)
}

func TestMerge(t *testing.T) {
	a := &BinTree{}
	a.InsLeft("left tree")

	b := &BinTree{}
	b.InsRight("right tree")

	c := Merge(a, b, "root")
	if c.root.Data != "root" || c.root.left != a.root || c.root.right != b.root {
		t.Error("合并失败")
	}
}

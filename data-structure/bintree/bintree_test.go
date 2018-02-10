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

func TestBinTreeRemove(t *testing.T) {
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

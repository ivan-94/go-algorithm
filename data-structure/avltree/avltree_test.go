package avltree

import "testing"

func Compare(a, b interface{}) int {
	_a := a.(int)
	_b := b.(int)
	if int(_a) > int(_b) {
		return 1
	} else if int(_a) == int(_b) {
		return 0
	}
	return -1
}

func TestInsert(t *testing.T) {
	tree := New(Compare)
	tree.Insert(43)
	tree.Insert(10)
	tree.Insert(50)
	if tree.len != 3 {
		t.Errorf("AVL插入异常, 长度应该为3, 实际上是: %d", tree.len)
	}
}

func Equal(n *Node, value int) bool {
	act := n.data.(int)
	return value == act
}

func TestLeftLeftRotate(t *testing.T) {
	tree := New(Compare)
	tree.Insert(35)

	// 35
	//   \
	//   48
	tree.Insert(48)
	if tree.root.right == nil || !Equal(tree.root.right, 48) {
		t.Error("插入异常, 应该插入右节点")
	}

	//    35
	//   /  \
	// 28   48
	tree.Insert(28)
	if tree.root.left == nil || !Equal(tree.root.left, 28) {
		t.Error("插入异常, 应该插入左节点")
	}

	//     35
	//    /  \
	//   28  48
	//  /  \
	// 20  30
	tree.Insert(20)
	tree.Insert(30)

	//      35(+2)
	//     /       \
	//   28(+1)     48
	//   /     \
	//  20(+1)  30
	//  /
	// 15
	tree.Insert(15)
	// expect
	//      28
	//     /   \
	//   20(+1)  35
	//   /       /  \
	//  15     30    48
	if !Equal(tree.root, 28) {
		t.Fatalf("翻转异常, root应该为%d, 实际上为 %v", 28, tree.root.Data())
	}

	if !Equal(tree.root.right, 35) {
		t.Fatalf("翻转异常, root.right应该为%d, 实际上为 %v", 25, tree.root.right.Data())
	}

	if !Equal(tree.root.left, 20) || tree.root.left.factor != avlLeftHeavy {
		t.Fatalf("翻转异常, root.left%d, 实际上为 %v", 25, tree.root.left.Data())
	}

	if !Equal(tree.root.right.left, 30) {
		t.Fatalf("翻转异常, root.right.left 应该为%d, 实际上为 %v", 25, tree.root.right.left.Data())
	}

	if !Equal(tree.root.left.left, 15) {
		t.Fatalf("翻转异常, root.left.left%d, 实际上为 %v", 25, tree.root.left.left.Data())
	}
}

func TestLeftRightRotate(t *testing.T) {
	// Insert
	//      35(+2)
	//     /       \
	//   28(-1)     48
	//   /     \
	//  20      30(-1)
	//            \
	//             32
	tree := New(Compare)
	tree.Insert(35)
	tree.Insert(48)
	tree.Insert(28)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(32)
	// Expect
	//      30
	//     /     \
	//   28(+1)   35
	//   /        /  \
	//  20      32    48
	if !Equal(tree.root, 30) {
		t.Fatalf("翻转异常, root %d, 实际上为 %v", 30, tree.root.Data())
	}

	if !Equal(tree.root.left, 28) || tree.root.left.factor != avlLeftHeavy {
		t.Fatalf("翻转异常, root %d, 实际上为 %v", 28, tree.root.left.Data())
	}

	if !Equal(tree.root.right.left, 32) {
		t.Fatalf("翻转异常, root %d, 实际上为 %v", 32, tree.root.right.left.Data())
	}
}

func TestRightRotate(t *testing.T) {
	// TODO:
}

func TestLookup(t *testing.T) {
	tree := New(Compare)
	tree.Insert(45)
	n, ok := tree.Lookup(45)
	if !ok || !Equal(n, 45) {
		t.Errorf("查找失败, 没有找到45")
	}

	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(35)
	tree.Insert(35)
	tree.Insert(37)

	n, ok = tree.Lookup(37)
	if tree.len != 6 || !ok || !Equal(n, 37) {
		t.Errorf("查找失败, 没有找到37")
	}

	tree.Remove(37)
	n, ok = tree.Lookup(37)
	if tree.len != 5 || ok || n != nil {
		t.Errorf("查找异常, 37已经隐藏")
	}

	tree.Clear()
	if tree.len != 0 || tree.root != nil {
		t.Errorf("树释放失败, 树的长度为: %d", tree.len)
	}
}

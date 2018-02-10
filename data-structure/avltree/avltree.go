// Package avltree 实现AVL实现的平衡搜索树
package avltree

const (
	avlBlanced    = 0
	avlLeftHeavy  = 1
	avlRightHeavy = -1
)

// Comparation 比较函数类型
type Comparation = func(a, b interface{}) int

// Node Avl 树节点
type Node struct {
	data        interface{}
	factor      int  // 平衡因子
	hidden      bool // 是否已删除
	left, right *Node
}

// Data 获取节点的数据
func (n *Node) Data() interface{} {
	return n.data
}

// AVLTree 代表AVL平衡树
type AVLTree struct {
	root    *Node
	len     int
	compare Comparation
}

// New 新建一个AVL平衡树
func New(compare Comparation) *AVLTree {
	return &AVLTree{
		compare: compare,
	}
}

// Insert 插入数据
func (t *AVLTree) Insert(data interface{}) {
	if t.root == nil {
		// 空树
		t.root = &Node{data: data}
		t.len++
	} else {
		t.insert(data, &t.root)
	}
}

// 插入一个元素到树种
func (t *AVLTree) insert(data interface{}, n **Node) (balance bool) {
	balance = true
	node := *n
	ndata := node.data
	cmpval := t.compare(data, ndata)

	if cmpval > 0 {
		// 插入右分支
		if node.right == nil {
			// 直接插入
			node.right = &Node{data: data}
			balance = false
		} else {
			// 向下递归
			balance = t.insert(data, &node.right)
		}

		if !balance {
			switch node.factor {
			case avlBlanced:
				// 右倾
				node.factor = avlRightHeavy
			case avlRightHeavy:
				// 右旋转
				t.rotateRight(n)
				balance = true
			case avlLeftHeavy:
				// 平衡了
				node.factor = avlBlanced
				balance = true
			}
		}

	} else if cmpval < 0 {
		// 插入左分支
		if node.left == nil {
			// 直接插入
			node.left = &Node{data: data}
			balance = false
		} else {
			balance = t.insert(data, &node.left)
		}

		if !balance {
			switch node.factor {
			case avlBlanced:
				node.factor = avlLeftHeavy
			case avlRightHeavy:
				node.factor = avlBlanced
				balance = true
			case avlLeftHeavy:
				t.rotateLeft(n)
				balance = true
			}
		}

	} else if node.hidden {
		// 相等， 判断是否当前节点是否已经删除， 如果已经删除， 则标记为未删除
		node.hidden = false
	}

	return
}

// 左旋转
func (t *AVLTree) rotateLeft(node **Node) {
	left := (*node).left
	if left.factor == avlLeftHeavy {
		// LL
		//       A(+2)                              B
		//      /    \                             /  \
		//     B(+1)  C                    =>     D(+1) A(+1)
		//    /    \                             /     /  \
		//   D(+1)  E                           X      E   C
		//   /
		//  X <- insert here
		(*node).left = left.right // A.left = B.right
		left.right = *node        // B.right = A
		(*node).factor = avlBlanced
		left.factor = avlBlanced
		*node = left
	} else {
		// LR
		//       A(+2)                              E
		//      /    \                             /  \
		//     B(+1)  C                    =>     B    A(-1)
		//    /    \                             / \   /   \
		//   D      E(+1)                       D   X  *    C
		//          / \
		//         X   *
		//         ^-insert here
		grandchild := left.right
		left.right = grandchild.left
		grandchild.left = left
		(*node).left = grandchild.right
		grandchild.right = *node

		switch grandchild.factor {
		case avlLeftHeavy:
			(*node).factor = avlRightHeavy
			left.factor = avlBlanced
		case avlBlanced:
			(*node).factor = avlBlanced
			left.factor = avlBlanced
		case avlRightHeavy:
			(*node).factor = avlBlanced
			left.factor = avlLeftHeavy
		}
		grandchild.factor = avlBlanced
		*node = grandchild
	}
}

// 右旋转
func (t *AVLTree) rotateRight(node **Node) {
	right := (*node).right
	if right.factor == avlRightHeavy {
		// RR
		//       A(-2)                               B
		//      /    \                             /   \
		//     C      B(-1)                =>     A     D(-1)
		//           / \                         / \      \
		//          E   D(-1)                   C   E       X
		//               \
		//                X <- insert here
		(*node).right = right.left
		right.left = *node
		(*node).factor = avlBlanced
		right.factor = avlBlanced
		*node = right
	} else {
		// RL
		//       A(-2)                               E
		//      /     \                             /   \
		//     C      B(+1)                =>     A      B(-1)
		//           /    \                      / \    /  \
		//          E(+1)   D                   C   X  *    D
		//         / \
		//        X   *
		//        ^-insert here
		grandchild := right.left
		right.left = grandchild.right
		grandchild.right = right
		(*node).right = grandchild.left
		grandchild.left = *node

		switch grandchild.factor {
		case avlLeftHeavy:
			(*node).factor = avlBlanced
			right.factor = avlRightHeavy
		case avlRightHeavy:
			(*node).factor = avlLeftHeavy
			right.factor = avlBlanced
		case avlBlanced:
			(*node).factor = avlBlanced
			right.factor = avlBlanced
		}
		grandchild.factor = avlBlanced
		*node = grandchild
	}
}

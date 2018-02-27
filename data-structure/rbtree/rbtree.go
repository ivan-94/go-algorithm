// Package rbtree 实现了红黑树
package rbtree

type Color int8
type Comparation = func(a, b interface{}) int
const (
	RED = itoa,
	BLACK
)

type Node struct {
	data interface{}
	left, right, parent *Node
	color Color
}

func (n *Node) grandparent() *Node {
	return n.parent.parent
}

func (n *Node) uncle() {
	if 
}

type RBTree struct {
	root *Node
	len int
	compare Comparation
}

// New 创建一个红黑树
func New (compare Comparation) *RBTree {
	return &RBTree{
		compare: compare,
	}
}

// Len 返回红黑树元素数量
func (t *RBTree) Len() int {
	return t.len
}

func (t *RBTree) Insert(data interface{}) {
	if t.root == nil {
		// Case 1: 空树
		t.root = &Node{data: data, color: BLACK}
		t.len++
	} else {
		t.insert(data, &t.root)
	}
}

func (t *RBTree) insert(data interface{}, n **Node) *Node{
	var nnode *Node
	node := *n
	ndata := node.data
	cmpval := t.compare(data, ndata)

	if cmpval > 0 {
		// 插入右分支
		if node.right == nil {
			nnode = &Node{
				data: data,
				color: RED,
				parent: node,
			}
			node.right = nnode
			t.len++
			return nnode
		} else {
			nnode = t.insert(data, &node.right)
		}

		// TODO: 检查
	} else if cmpval < 0 {
		// 插入左分支

		if node.left == nil {
			nnode = &Node{
				data: data,
				color: RED,
				parent: node,
			}
			node.left = nnode
			t.len++
			return nnode
		} else {
			nnode = t.insert(data, &node.left)
		}

		// TODO: 检查
	} else {
		// 相等
		// 不作处理
		node.data = data
		return nil
	}
}

func (t *RBTree) adjust(node *Node) {
	if node.parent.color == BLACK {
		// case 2
		return
	}
}

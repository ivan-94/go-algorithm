// Package bintree 实现了二叉树数据结构
package bintree

import "errors"

// 错误类型定义
var (
	ErrInsConflit = errors.New("node already existed")
	ErrRemoved    = errors.New("node removed")
	ErrEmpty      = errors.New("node empty")
)

// Node 代表二叉树的节点
type Node struct {
	Data        interface{}
	left, right *Node
}

// Left 获取左节点
func (n *Node) Left() *Node {
	return n.left
}

// Right 获取右节点
func (n *Node) Right() *Node {
	return n.left
}

// IsLeaf 判断是否是叶子节点
func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// BinTree 代表二叉树类型
type BinTree struct {
	root *Node
	len  int
}

// Root 获取根节点
func (t *BinTree) Root() *Node {
	return t.root
}

func (t *BinTree) insert(dir string, data interface{}, node ...*Node) (*Node, error) {
	var pos **Node

	if len(node) > 0 {
		// 插入指定节点
		n := node[0]
		if (dir == "left" && n.left != nil) || (dir == "right" && n.right != nil) {
			return nil, ErrInsConflit
		}
		if dir == "left" {
			pos = &n.left
		} else {
			pos = &n.right
		}
	} else {
		// 插入到根节点
		if t.root != nil {
			return t.insert(dir, data, t.root)
		}
		pos = &t.root
	}

	nn := &Node{Data: data}
	*pos = nn
	t.len++
	return nn, nil

}

func (t *BinTree) remove(dir string, node ...*Node) (*Node, error) {
	var removed *Node
	if len(node) > 0 {
		el := node[0]
		if el == nil {
			return nil, ErrEmpty
		}

		if (dir == "left" && el.left == nil) || (dir == "right" && el.right == nil) {
			return nil, ErrRemoved
		}
		if dir == "left" {
			removed = el.left
			el.left = nil
		} else {
			removed = el.right
			el.right = nil
		}
	} else {
		// remove from root
		if t.root != nil && t.root.left == nil && t.root.right == nil {
			removed = t.root
			t.root = nil
		} else {
			return t.remove(dir, t.root)
		}
	}
	t.len--
	return removed, nil
}

// RmLeft 移除左节点
func (t *BinTree) RmLeft(node ...*Node) (*Node, error) {
	return t.remove("left", node...)
}

// RmRight 移除右节点
func (t *BinTree) RmRight(node ...*Node) (*Node, error) {
	return t.remove("right", node...)
}

// InsLeft 插入左节点, 如果没有给定节点, 则插入到根节点
func (t *BinTree) InsLeft(data interface{}, node ...*Node) (*Node, error) {
	return t.insert("left", data, node...)
}

// InsRight 插入右节点, 如果没有给定节点, 则插入到根节点
func (t *BinTree) InsRight(data interface{}, node ...*Node) (*Node, error) {
	return t.insert("right", data, node...)
}

// Len 获取节点树
func (t *BinTree) Len() int {
	return t.len
}

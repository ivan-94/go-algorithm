package queue

import (
	"github.com/carney520/go-algorithm/data-structure/list"
)

// Queue 表示队列类型
type Queue struct {
	list *list.List
}

// Enqueue 插入队列
func (q *Queue) Enqueue(data interface{}) {
	q.list.Append(data)
}

// Dequeue 出队列
func (q *Queue) Dequeue() interface{} {
	head := q.list.Head()
	if head != nil {
		q.list.Remove(head)
		return head.Data
	}
	return nil
}

// Peek 获取队头数据, 但不出队列
func (q *Queue) Peek() interface{} {
	head := q.list.Head()
	if head != nil {
		return head.Data
	}
	return nil
}

// Len 队列长度
func (q *Queue) Len() int {
	return q.list.Len()
}

// New 创建一个队列
func New() *Queue {
	return &Queue{
		list: list.New(),
	}
}

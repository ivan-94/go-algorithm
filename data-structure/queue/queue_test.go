package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Len() != 3 || q.Dequeue() != 1 || q.Dequeue() != 2 || q.Dequeue() != 3 || q.Len() != 0 {
		t.Errorf("队列弹出异常， 没有按照先进先出的规则返回数据")
	}

	q.Enqueue(1)
	if q.Peek() != 1 || q.Dequeue() != 1 || q.Len() != 0 {
		t.Errorf("队列弹出异常，Peek返回数据异常")
	}

	if q.Dequeue() != nil {
		t.Errorf("队列弹出异常，空队列弹出应该为nil")
	}
}

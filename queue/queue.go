package queue

import (
	"container/list"
)

var _ = list.New()

// Queue 通用队列
type Queue interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
}

type queue struct {
	list *list.List
}

// New 实例化Queue
func New() Queue {
	return &queue{
		list: list.New(),
	}
}

// Push 入队
func (q *queue) Push(v interface{}) {
	q.list.PushBack(v)
}

// Pop 出队
func (q *queue) Pop() interface{} {
	v := q.list.Front()
	if v == nil {
		return nil
	}

	return q.list.Remove(v)
}

// 获取队头数据,但是不移除
func (q *queue) Peek() interface{} {
	v := q.list.Front()
	if v == nil {
		return nil
	}
	return v.Value
}

// Len 队列长度
func (q *queue) Len() int {
	return q.list.Len()
}

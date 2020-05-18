package queue

import (
	"container/list"
	"fmt"
)

//Queue ...
type Queue struct {
	list *list.List
}

//New new Queue()
func New() *Queue {
	return &Queue{list: &list.List{}}
}

//Push add item
func (s *Queue) Push(val int) {
	s.list.PushBack(val)
}

// Pop get and remove last item
func (s *Queue) Pop() {
	if s.Empty() {
		return
	}
	e := s.list.Front()
	fmt.Println(s.list.Remove(e))
}

//Empty ...
func (s *Queue) Empty() bool {
	if s.list.Len() == 0 {
		return true
	}
	return false
}

//Front ...
func (s *Queue) Front() interface{} {
	return s.list.Front().Value
}

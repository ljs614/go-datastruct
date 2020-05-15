package stack

import (
	"container/list"
	"fmt"
)

//Stack ...
type Stack struct {
	list *list.List
}

//New new Stack()
func New() *Stack {
	return &Stack{list: &list.List{}}
}

//Push add item
func (s *Stack) Push(val int) {
	s.list.PushBack(val)
}

// Pop get and remove last item
func (s *Stack) Pop() {
	if s.Empty() {
		return
	}
	e := s.list.Back()
	fmt.Println(s.list.Remove(e))
}

//Empty ...
func (s *Stack) Empty() bool {
	if s.list.Len() == 0 {
		return true
	}
	return false
}

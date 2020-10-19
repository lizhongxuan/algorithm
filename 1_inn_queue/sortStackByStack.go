//题目:用一个栈实现另一个栈的排序
package main

func (s *Stack)stackSort()  {
	s2 := NewStack()

	for !s.isEmpty() {
		v := s.Pop()
		for !s2.isEmpty() && s2.top.value.(int) > v.(int) {
			v2 := s2.Pop()
			s.Push(v2)
		}
		s2.Push(v)
	}

	for !s2.isEmpty()  {
		v2 := s2.Pop()
		s.Push(v2)
	}

}
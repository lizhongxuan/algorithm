//题目:如何仅用递归函数和栈操作逆序一个栈

package main

func (s *Stack)reverseOrder(){
	if s.top == nil{
		return
	}
	v := s.getAndRemoveLastElement()



	s.reverseOrder()
	s.Push(v)
}

func (s *Stack)getAndRemoveLastElement()interface{}{
	if s.botton == nil {
		return nil
	}

	v := s.botton.value
	if s.botton.preNode == nil{
		s.top = nil
		s.botton = nil
		return v
	}

	s.botton = s.botton.preNode
	return v
}

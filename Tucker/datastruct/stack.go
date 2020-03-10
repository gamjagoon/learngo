package datastruct

//Stack : struct
type Stack struct {
	ll *Linklist
}

//NewStack : make func
func NewStack() *Stack {
	return &Stack{ll: &Linklist{}}
}

//Push : addval
func (s *Stack) Push(val int) {
	s.ll.AddTail(val)
}

//Top : return val
func (s *Stack) Top() int {
	ret := s.ll.Back()
	return ret
}

//Pop : remove top
func (s *Stack) Pop() {
	s.ll.Removelast()
}

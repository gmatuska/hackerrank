package stack

type MyStack struct {
	items []rune
}

func (s *MyStack) Push(item rune) *MyStack {
	s.items = append(s.items, item)
	return s
}

func (s *MyStack) MirrorPush(ms *MyStack, item rune) (*MyStack, *MyStack, bool) {
	var mirror rune
	switch item {
	case '[':
		mirror = ']'
		break
	case '(':
		mirror = ')'
		break
	case '{':
		mirror = '}'
		break
	default:
		mirror = 0
		break
	}
	if mirror == 0 {
		return s, ms, false
	} else {
		return s.Push(item), ms.Push(mirror),true
	}
}

func (s *MyStack) Pop() rune {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *MyStack) Empty()string{
	var empty []rune
	for range s.items{
		empty = append(empty,s.Pop())
	}
	return string(empty)
}

func NewStack() *MyStack {
	return &MyStack{}
}

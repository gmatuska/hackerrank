package stack

import "testing"

func TestLastInFirstOut(t *testing.T) {
	s := NewStack()
	s.Push('r')
	s.Push('u')
	s.Push('n')
	s.Push('e')
	item := s.Pop()
	if item != 'e' {
		t.Errorf("Expected e. Actual was %v", item)
	}
}

func TestUnbalancedBracket(t *testing.T) {
	s := NewStack()
	mirrorS := NewStack()
	s, mirrorS, ok := s.MirrorPush(mirrorS, '[')
	if !ok {
		//compare rest of string with mirror stack
		mirror := mirrorS.Empty()
		if mirror != "]" {
			t.Errorf("Expected ']'. Actual was %v", mirrorS.Empty())
		}
	}
}

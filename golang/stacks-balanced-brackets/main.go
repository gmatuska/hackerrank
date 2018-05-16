package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tests_str, err := ReadLine(reader)
	Check(err)
	tests, err := strconv.Atoi(tests_str)
	var results []string
	for i := 0; i < tests; i++ {
		br_str, err := ReadLine(reader)
		Check(err)
		s := NewStack()
		//empty and compare with remaining test string
		if s.IsBalanced(br_str){
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	for _, i := range results {
		fmt.Println(i)
	}
}

func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

type MyStack struct {
	items []rune
}

func (s *MyStack) IsBalanced(str string)bool{
	runes := []rune(str)
	for _,c := range runes{
		switch c{
		case '[':
			s.Push(']')
			break
		case '{':
			s.Push('}')
			break
		case '(':
			s.Push(')')
			break
		default:
			if s.Empty() || c != s.Top(){
				return false
			}
			_ = s.Pop()
			break
		}
	}
	return s.Empty()
}
func (s *MyStack) Push(item rune) *MyStack {
	s.items = append(s.items, item)
	return s
}

func (s *MyStack) Pop() rune {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *MyStack) Top() rune {
	return s.items[len(s.items)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.items) == 0
}

func NewStack() *MyStack {
	return &MyStack{}
}

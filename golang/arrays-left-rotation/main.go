package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"hackerrank/arrays-left-rotation/queue"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n_string, err := ReadLn(reader)
	Check(err)
	tokens, err := Slice_Atoi(strings.Split(n_string, " "))
	n := tokens[0]
	k := tokens[1]
	n_string, err = ReadLn(reader)
	tokens, err = Slice_Atoi(strings.Split(n_string, " "))
	q := queue.New(n)
	//enqueue
	for _, i := range tokens {
		q.Enqueue(i)
	}
	//dequeue and enqueue
	for j := 0; j < k; j++ {
		q.Enqueue(q.Dequeue())
	}
	for j := 0; j < n; j++ {
		fmt.Print(q.Dequeue(), " ")
	}
}

func ReadLn(r *bufio.Reader) (string, error) {
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

func Slice_Atoi(strArr []string) ([]int, error) {
	// NOTE:  Read Arr as Slice as you like
	var str string // O
	var i int      // O
	var err error  // O

	iArr := make([]int, 0, len(strArr))
	for _, str = range strArr {
		i, err = strconv.Atoi(str)
		if err != nil {
			return nil, err // O
		}
		iArr = append(iArr, i)
	}
	return iArr, nil
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

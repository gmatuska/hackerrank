package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	p_string, _ := ReadLine(reader)
	p, err := strconv.Atoi(p_string)
	check(err)
	q_string, _ := ReadLine(reader)
	q, err := strconv.Atoi(q_string)
	check(err)
	var kaprekars []int
	for i := q; i >= p; i-- {
		//get the length of i
		if Contains(kaprekars,i) {
			continue
		}
		d := len(strconv.Itoa(i))
		sq := int64(i * i)
		sqAr := strings.SplitAfter(strconv.FormatInt(sq,10),"")
		sqAr0 := sqAr[:len(sqAr)-d]
		sq0 := 0
		sq1 := 0
		sq0,err = strconv.Atoi(strings.Join(sqAr0,""))
		sqAr1 := sqAr[len(sqAr)-d:]
		sq1,err = strconv.Atoi(strings.Join(sqAr1,""))
		if sq0 + sq1 == i {
			kaprekars = append(kaprekars,i)
		}
	}
	if len(kaprekars) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		for j := len(kaprekars)-1; j >= 0; j-- {
			fmt.Print(kaprekars[j])
			if j > 0 {
				fmt.Print(" ")
			}
		}
	}
}

func Contains(slice []int, item int) bool {
	set := make(map[int]struct{},len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

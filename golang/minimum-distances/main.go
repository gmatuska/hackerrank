package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	n_string,_ := Readln(reader)
	n,err := strconv.Atoi(n_string)
	check(err)
	a_string,_ := Readln(reader)
	a,err := Slice_Atoi(strings.Split(a_string," "))
	used := make(map[int]int)
	maxInt := 9223372036854775807
	min := maxInt
	for i := 0; i < n; i++ {
		if Contains(used,a[i]) {
			used[a[i]] = i - used[a[i]]
			//find all indices of a[i] in a
			if used[a[i]] < min {
				min = used[a[i]]
			}
		} else {
			used[a[i]] = i
		}
	}
	if min == maxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(min)
	}
}
func Contains(set map[int]int,item int) bool {
	_,ok := set[item]
	return ok
}
func Readln(r *bufio.Reader) (string, error) {
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
func check(e error) {
	if e != nil {
		panic(e)
	}
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

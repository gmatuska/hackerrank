package main

import (
	"bufio"
	"os"
	"sort"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t_string, _ := Readln(reader)
	t,err := strconv.Atoi(t_string)
	check(err)
	for i := 0; i < t; i++ {
		w, _ := Readln(reader)

		wAr := strings.SplitAfter(w,"")
		success := next(sort.StringSlice(wAr))
		if success {
			fmt.Println(strings.Join(wAr,""))
		} else {
			fmt.Println("no answer")
		}
	}
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
func next(data sort.Interface) bool {
	var k, l int
	for k = data.Len() - 2; ; k-- {
		if k < 0 {
			return false
		}
		if data.Less(k,k+1) {
			break
		}
	}
	for l = data.Len() - 1; !data.Less(k,l); l-- {
	}
	data.Swap(k,l)
	for i,j := k+1, data.Len()-1;i < j;i++ {
		data.Swap(i,j)
		j--
	}
	return true
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

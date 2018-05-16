package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nd_string, _ := Readln(reader)
	nd,err := Slice_Atoi(strings.Split(nd_string," "))
	check(err)
	d := nd[1]
	a_string, _ := Readln(reader)
	a,err := Slice_Atoi(strings.Split(a_string," "))
	aMap := Convert(a)
	beautifuls := 0
	for i := range a {
		blah := a[i]+d
		blah2 := a[i]+(2*d)
		if Contains(aMap,blah) && Contains(aMap,blah2) {
			beautifuls++
		}
	}
	fmt.Println(beautifuls)

}
func Contains(set map[int]struct{},item int) bool {
	_, ok := set[item]
	return ok
}
func Convert(slice []int) map[int]struct{} {
	set := make(map[int]struct{},len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	return set
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

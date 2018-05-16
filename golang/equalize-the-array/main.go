package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	slice_string, _ := Readln(reader)
	slice,err := Slice_Atoi(strings.Split(slice_string," "))
	if err != nil {
		panic(err)
	}
	sort.Ints(slice)
	theNumber := 0
	count := 0
	largestCount := 0
	for i := range slice {
		if slice[i] != theNumber {
			if largestCount < count {
				largestCount = count
			}
			count = 1
			theNumber = slice[i]
		} else {
			count++
		}
	}
	if largestCount < count {
		largestCount = count
	}
	fmt.Println(n-largestCount)
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

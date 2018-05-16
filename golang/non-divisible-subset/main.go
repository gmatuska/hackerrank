package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nk_string, _ := Readln(reader)
	nk, err := Slice_Atoi(strings.Split(nk_string, " "))
	if err != nil {
		panic(err)
	}
	k := nk[1]
	//load in the set S
	s_string, _ := Readln(reader)
	numbers,err := Slice_Atoi(strings.Split(s_string," "))
	if err != nil {
		panic(err)
	}

	//create map for integers
	counts := make([]int,k,k)
	for number := range numbers {
		index := numbers[number] % k
		counts[index] += 1
	}
	count := Min(counts[0],1)
	max := (k/2)+1
	for i := 1; i < max; i++ {
		if i != k - i {
			count += Max(counts[i],counts[k-i])
		}
	}
	if k % 2 == 0 {
		count += 1
	}
	fmt.Println(count)
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

func Min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

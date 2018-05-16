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
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	clouds_string, _ := Readln(reader)
	clouds,err := Slice_Atoi(strings.Split(clouds_string," "))
	if err != nil {
		panic(err)
	}
	//cloud[i] can only go to cloud[i+1] or cloud[i+2]
	//so to get to cloud[n-1], we would prefer to jump to cloud[i+2]
	count := -1
	for i := 0; i < n; i++ {
		if i + 2 < n && clouds[i+2] == 0 {
			i++
		}
		count++
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


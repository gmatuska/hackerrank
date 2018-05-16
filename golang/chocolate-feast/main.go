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
	t_string,_ := ReadLine(reader)
	t,err := strconv.Atoi(t_string)
	Check(err)
	results := make([]int,t,t)
	for i := 0; i < t; i++ {
		ncm_string, _ := ReadLine(reader)
		ncmAr,err := Slice_Atoi(strings.Split(ncm_string," "))
		Check(err)
		n := ncmAr[0]
		c := ncmAr[1]
		m := ncmAr[2]
		count := n / c
		wrappers := count
		for wrappers >= m {
			wrappers = wrappers - m
			count++
			wrappers++
		}
		results[i] = count
	}
	for i := range results {
		fmt.Println(results[i])
	}
}

func ReadLine(reader *bufio.Reader) (string,error) {
	var (
		isPrefix bool	= true
		err error	= nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line,isPrefix,err = reader.ReadLine()
		ln = append(ln,line...)
	}
	return string(ln),err
}

func Check(err error) {
	if err != nil {
		panic(err)
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

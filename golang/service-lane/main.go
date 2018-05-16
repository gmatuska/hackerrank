package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nt_string, err := ReadLine(reader)
	nt,err := Slice_Atoi(strings.Split(nt_string," "))
	Check(err)
	n := nt[0]
	t := nt[1]
	results := make([]int,t,t)
	width_string, err := ReadLine(reader)
	width,err := Slice_Atoi(strings.Split(width_string," "))
	Check(err)
	for a := 0; a < t; a++ {
		ij_string, _ := ReadLine(reader)
		ij,err := Slice_Atoi(strings.Split(ij_string," "))
		i := ij[0]
		j := ij[1]
		Check(err)
		temp := make([]int,n,n)
		copy(temp,width)
		temp = temp[i:j+1]
		results[a] = Min(temp)
	}
	for i := range results {
		fmt.Println(results[i])
	}

}
func Min(iArr []int) int {
	sort.Ints(iArr)
	min := iArr[0]
	if min > 3 {
		return 3
	} else {
		return min
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

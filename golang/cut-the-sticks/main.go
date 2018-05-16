package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"fmt"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	//load the number of sticks (N)
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	//load the sticks
	sticks_string, _:= Readln(reader)
	var sticks []int
	for a := 0; a < n; a++ {
		sticks, err = Slice_Atoi(strings.Split(sticks_string, " "))
		if err != nil {
			panic(err)
		}
	}

	//loop until no sticks are left
	fmt.Println(len(sticks))
	for len(sticks) > 0 {
		//the value of the cut-operation is the minimum value
		//so sort sticks and grab the first value in the slice
		sort.Ints(sticks)
		cut := sticks[0]
		k := 0
		deleted := 0
		for _, j := range sticks {
			j = j - cut
			sticks[k] = j
			k++
			if  j <= 0 {
				//need to delete the element
				deleted++
			}
		}
		sticks = sticks[deleted:]
		if len(sticks) > 0 {
			fmt.Println(len(sticks))
		}
	}
}

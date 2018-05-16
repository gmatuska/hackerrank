package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main(){
	//so that we can compare them with the given square
	/*
	8 1 6
	3 5 7
	4 9 2
	*****
	6 1 8
	7 5 2
	2 9 4
	*****
	4 9 2
	3 5 7
	8 1 6
	*****
	2 9 4
	7 5 3
	6 1 8
	*****
	8 3 4
	1 5 9
	6 7 2
	*****
	4 3 8
	9 5 1
	2 7 6
	*****
	6 7 2
	1 5 9
	8 3 4
	*****
	2 7 6
	9 5 1
	4 3 8
	 */
	//above are all the 3x3 magic squares possible
	//read in the squares, and evaluate
	//but first, load in the magic squares as slices

	pre := [][][] int {
		{{2,7,6},{9,5,1},{4,3,8}},
		{{6,7,2},{1,5,9},{8,3,4}},
		{{4,3,8},{9,5,1},{2,7,6}},
		{{8,3,4},{1,5,9},{6,7,2}},
		{{2,9,4},{7,5,3},{6,1,8}},
		{{4,9,2},{3,5,7},{8,1,6}},
		{{6,1,8},{7,5,3},{2,9,4}},
		{{8,1,6},{3,5,7},{4,9,2}},
	}

	reader := bufio.NewReader(os.Stdin)
	magic_square := make([][]int,3,3)

	for a := 0; a < 3; a++ {
		n_string, _ := Readln(reader)
		n, err := Slice_Atoi(strings.Split(n_string, " "))
		if err != nil {
			panic(err)
		}
		magic_square[a] = n
	}
	min_cost := 81
	var crt_cost int
	for k := 0; k < 8; k++ {
		crt_cost = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				temp := magic_square[i][j] - pre[k][i][j]
				if temp < 0 {
					temp *= -1
				}
				crt_cost += temp
			}
		}
		if crt_cost < min_cost {
			min_cost = crt_cost
		}
	}
	fmt.Println(min_cost)
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

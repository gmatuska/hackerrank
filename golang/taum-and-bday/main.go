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
	t_string, _ := Readln(reader)
	t, err := strconv.Atoi(t_string)
	check(err)
	sums := make([]int,t,t)
	for i := 0; i < t; i++ {
		gifts_string, _ := Readln(reader)
		gifts,err := Slice_Atoi(strings.Split(gifts_string," "))
		check(err)
		black := gifts[0]
		white := gifts[1]
		values_string, _ := Readln(reader)
		values,err := Slice_Atoi(strings.Split(values_string," "))
		x := values[0]
		y := values[1]
		z := values[2]
		//if z is less than either x or y, then you can convert to the lowest of the 2
		switch {
		case x < y:
			if x + z < y {
				y = x + z
			}
			break
		case y < x:
			if y + z < x {
				x = y + z
			}
			break
		default:
			break
		}
		sums[i] = (black * x) + (white * y)
	}
	for i := range sums {
		fmt.Println(sums[i])
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

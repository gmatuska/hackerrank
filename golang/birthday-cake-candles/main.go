package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
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
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	fmt.Println("Enter the number of candles:")
	reader := bufio.NewReader(os.Stdin)
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	candles_string, _ := Readln(reader)
	candles, err := Slice_Atoi(strings.Split(candles_string, " "))
	if err != nil {
		panic(err)
	}
	largest := candles[0]
	largest_count := 1
	for i := 1; i < n; i++ {
		switch {
		case candles[i] > largest:
			largest = candles[i]
			largest_count = 1
		case candles[i] == largest:
			largest_count++
		}
	}
	fmt.Println(largest_count)
}

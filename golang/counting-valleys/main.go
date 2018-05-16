package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

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
	elevation := 0
	valleys := 0
	reader := bufio.NewReader(os.Stdin)
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	steps, _ := Readln(reader)
	//iterate through each character in steps
	for i := 0; i < n; i++ {
		if steps[i] == 'D' {
			elevation--
		} else {
			elevation++
			if elevation == 0 {
				valleys++
			}
		}
	}
	fmt.Println(valleys)
}
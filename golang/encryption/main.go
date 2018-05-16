package main

import (
	"math"
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text,err := ReadLine(reader)
	check(err)
	a := math.Sqrt(float64(len(text)))
//	floor := math.Floor(a)
	ceil := math.Ceil(a)
	count := 0
	for i := 0; i < int(ceil); i++ {
		j := i
		for j < len(text) {
			fmt.Print(string(text[j]))
			j += int(ceil)
			count++
		}
		fmt.Print(" ")
	}
}
func ReadLine(r *bufio.Reader) (string, error) {
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

package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := Readln(reader)
	//load in the set S
	n_string, _ := Readln(reader)
	n,err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Count(s,"a") * (n / len(s)) + strings.Count(s[:n % len(s)],"a"))
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

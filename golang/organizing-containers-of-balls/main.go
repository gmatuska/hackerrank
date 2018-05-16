package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"sort"
	"reflect"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	q_string, _ := Readln(reader)
	q, err := strconv.Atoi(q_string)
	results := make([]string,q,q)
	check(err)
	for i := 0; i < q; i++ {
		n_string, _ := Readln(reader)
		n,err := strconv.Atoi(n_string)
		//create ball type array
		balls := make([]int,n,n)
		containers := make([]int,n,n)
		check(err)
		for j := 0; j < n; j++ {
			container_string, _ := Readln(reader)
			container,err := Slice_Atoi(strings.Split(container_string," "))
			check(err)
			for k := range container {
				containers[j] += container[k]
			}
			for k := 0; k < n; k++ {
				balls[k] += container[k]
			}
		}
		sort.Ints(balls)
		sort.Ints(containers)
		if reflect.DeepEqual(balls,containers) {
			results[i] = "Possible"
		} else {
			results[i] = "Impossible"
		}
	}
	for b := range results {
		fmt.Println(results[b])
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

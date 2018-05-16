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
	nm_string, _ := Readln(reader)
	nm, err := Slice_Atoi(strings.Split(nm_string," "))
	check(err)
	n := nm[0]
	m := nm[1]
	members := make([][]int,n)
//	var members [n][m]int
	b := 0
	for a := 0; a < n; a++ {
		topics_string, _ := Readln(reader)
//		topics_string = Reverse(topics_string)
		b = 0
		members[a] = make([]int,m)
		for _, r := range topics_string {
			members[a][b],err = strconv.Atoi(string(r))
			check(err)
			b++
		}
	}
	//go through each combo of 2 members: keep track of the max number of topics
	temp := make([]int,m,m)
	maxTopics := 0
	maxTopicsCombos := 0
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			for k := 0; k < m; k++ {
				temp[k] = members[i][k] | members[j][k]
			}
			tempCount := Count1s(temp)
			//compare the 1s count to the max topics value
			switch {
			case tempCount == maxTopics:
				maxTopicsCombos++
				break
			case tempCount > maxTopics:
				maxTopics = tempCount
				maxTopicsCombos = 1
				break
			default:
				break
			}
		}
	}
	fmt.Println(maxTopics)
	fmt.Println(maxTopicsCombos)
}
func Count1s(intArr []int) int {
	count := 0
	for i := range intArr {
		if intArr[i] == 1 {
			count++
		}
	}
	return count
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

func Min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return a * -1
	}
	return a
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SliceAtoI(strArr []string) ([]int, error) {
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

func ReadLn(r *bufio.Reader) (string, error) {
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

func main(){
	reader := bufio.NewReader(os.Stdin)

	firstLine_string, _ := ReadLn(reader)
	firstLine, err := SliceAtoI(strings.Split(firstLine_string, " "))
	if err != nil {
		panic(err)
	}
	s := firstLine[0]
	n := firstLine[1]
	m := firstLine[2]

	n_string, _ := ReadLn(reader)
	n_slice,err := SliceAtoI(strings.Split(n_string," "))
	if err != nil {
		panic(err)
	}
	m_string, _ := ReadLn(reader)
	m_slice,err := SliceAtoI(strings.Split(m_string," "))
	if err != nil {
		panic(err)
	}

	//sort keyboards in ascending order
	sort.Ints(n_slice)
	//sort USBs in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(m_slice)))
	//set default difference
	total := -1

	for i:=0;i < n;i++ {
		for j:=m-1;j >=0;j-- {
			temp := n_slice[i] + m_slice[j]
			switch {
			case temp > s: continue
			case temp == s: {
				total = temp
				continue
			}
			case s - temp < s - total: total = temp
			}
		}
		if total == s {
			break
		}
	}
	fmt.Println(total)
}

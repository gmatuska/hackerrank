package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nk_string,err := ReadLine(reader)
	Check(err)
	nk,err := Slice_Atoi(strings.Split(nk_string," "))
	t_string,err := ReadLine(reader)
	Check(err)
	t,err := Slice_Atoi(strings.Split(t_string," "))
	specials := 0
	k := nk[1]
	pages := 0
	for i := 1; i <= nk[0]; i++ {
		problems := make([]int,k)
		pagesPerChapter := t[i-1] / k
		for j := 0; j < pagesPerChapter; j++ {
			for l := 1; l <= k; l++ {
				problems[l-1] = k * j + l
			}
			if Contains(problems,pages + j + 1) {
				specials++
			}
		}
		if t[i-1] % k > 0 {
			remainderProblems := make([]int,t[i-1] % k)
			pagesPerChapter++
			begin := (pagesPerChapter-1) * k + 1
			end := begin + (t[i-1] % k)
			count := 0
			for a := begin; a < end; a++ {
				remainderProblems[count] = a
				count++
			}
			if Contains(remainderProblems,pages+pagesPerChapter){
				specials++
			}
		}
		pages += pagesPerChapter
	}
	fmt.Println(specials)
}
func Contains(arr []int,item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
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

func ReadLine(reader *bufio.Reader) (string,error) {
	var (
		isPrefix bool	= true
		err error	= nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line,isPrefix,err = reader.ReadLine()
		ln = append(ln,line...)
	}
	return string(ln),err
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}


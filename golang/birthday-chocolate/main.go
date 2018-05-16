package birthday_chocolate

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

func main(){
	reader := bufio.NewReader(os.Stdin)
	n_string, _ := Readln(reader)
	n, err := strconv.Atoi(n_string)
	if err != nil {
		panic(err)
	}
	squares_string, _ := Readln(reader)
	squares, err := Slice_Atoi(strings.Split(squares_string, " "))
	if err != nil {
		panic(err)
	}
	birth_string, _:= Readln(reader)
	birth,err := Slice_Atoi(strings.Split(birth_string," "))
	if err != nil {
		panic(err)
	}
	day := birth[0]
	month := birth[1]
	count := 0
	startIndex := 0
	for startIndex + month <= n {
		sum := 0
		for j := 0; j < month; j++ {
			sum += squares[startIndex+j]
		}
		if sum == day{
			count++
		}
		startIndex++
	}
	fmt.Println(count)
}

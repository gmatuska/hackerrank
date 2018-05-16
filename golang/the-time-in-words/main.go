package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	hour_string, err := ReadLine(reader)
	check(err)
	hour,err := strconv.Atoi(hour_string)
	check(err)
	mins_string, err := ReadLine(reader)
	mins,err := strconv.Atoi(mins_string)
	fmt.Println(ConvertTimeToWords(hour,mins))
}
func ConvertTimeToWords(hour,min int) string {
	dict := map[int]string{
		0: "o' clock",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "half past",
	}
	wordtime := ""
	switch {
	case min == 0:
		wordtime = dict[hour] + " " + dict[min]
		break
	case min == 15:
		wordtime = dict[min] + " past " + dict[hour]
		break
	case min == 30:
		wordtime = dict[30] + " " + dict[hour]
		break
	case min == 45:
		if hour == 12 {
			hour = 0
		}
		wordtime = dict[15] + " to " +dict[hour+1]
		break
	case min < 30:
		if min > 20 {
			mod := min % 10
			wordtime = "twenty " + dict[mod] + " minutes past " + dict[hour]
		} else {
			wordtime = dict[min] + " minutes past " + dict[hour]
		}
		break
	default:
		min = 60 - min
		if hour == 12 {
			hour = 0
		}
		if min > 20 {
			mod := min % 10
			wordtime = "twenty " + dict[mod] + " minutes to " + dict[hour+1]
		} else {
			wordtime = dict[min] +  " minutes to " + dict[hour+1]
		}
		break
	}
	return wordtime
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
func check(err error) {
	if err != nil {
		panic(err)
	}
}
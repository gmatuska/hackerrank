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
	nk_string, _ := Readln(reader)
	nk, err := Slice_Atoi(strings.Split(nk_string," "))
	check(err)
	n := nk[0]
	k := nk[1]
	queen_string, _ := Readln(reader)
	queen,err := Slice_Atoi(strings.Split(queen_string," "))
	if err != nil {
		panic(err)
	}
	xq := queen[1]
	yq := queen[0]
	dist_slice := make([]int,8,8)
	dist_slice[0] = n - yq
	dist_slice[2] = n - xq
	dist_slice[4] = yq - 1
	dist_slice[6] = xq - 1
	//for diagonals, take the min distance from the boundaries to the queen for x & y
	dist_slice[1] = Min(dist_slice[0],dist_slice[2])
	dist_slice[3] = Min(dist_slice[2],dist_slice[4])
	dist_slice[5] = Min(dist_slice[4],dist_slice[6])
	dist_slice[7] = Min(dist_slice[6],dist_slice[0])
	for i := 0; i < k; i++ {
		obstacle_string, _ := Readln(reader)
		obstacle,err := Slice_Atoi(strings.Split(obstacle_string," "))
		if err != nil {
			panic(err)
		}
		x := obstacle[1]
		y := obstacle[0]
		if x == xq && y == yq {
			continue
		}
		//is obstacle in the path of the queen?
		switch {
		case x == xq:
			if y > yq {
				dist_slice[0] = Min(dist_slice[0],Abs(y-yq)-1)
			} else {
				dist_slice[4] = Min(dist_slice[4],Abs(y-yq)-1)
			}
			break
		case y == yq:
			if x > xq {
				dist_slice[2] = Min(dist_slice[2],Abs(x-xq)-1)
			} else {
				dist_slice[6] = Min(dist_slice[6],Abs(x-xq)-1)
			}
			break
		case Abs(x-xq) == Abs(y-yq):
			if x > xq && y > yq {
				dist_slice[1] = Min(dist_slice[1],Min(Abs(x-xq)-1,Abs(y-yq)-1))
			} else if x > xq && y < yq {
				dist_slice[3] = Min(dist_slice[3],Min(Abs(x-xq)-1,Abs(y-yq)-1))
			} else if x < xq && y > yq {
				dist_slice[7] = Min(dist_slice[7],Min(Abs(x-xq)-1,Abs(y-yq)-1))
			} else if x < xq && y < yq {
				dist_slice[5] = Min(dist_slice[5],Min(Abs(x-xq)-1,Abs(y-yq)-1))
			}
		}
	}
	sum := 0
	for i := range dist_slice {
		sum += dist_slice[i]
	}
	fmt.Println(sum)
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
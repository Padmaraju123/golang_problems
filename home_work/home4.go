package main

import (
	"fmt"
	"strconv"
)

func countt(num string, c int) ([]int,int) {
	le := len(num)
	num1, _ := strconv.Atoi(num)
	var out = make([]int, le)
	for num1 > 0 {
		re := num1 % 10
		num1 = num1/10
		out[le-1] = re
		le--
		c++
	}

	return out,c

}

func main() {
	var num string
	var c int
	fmt.Scanln(&num)
	sli,count := countt(num, c)
	fmt.Println(sli,count)

}

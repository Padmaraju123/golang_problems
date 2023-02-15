package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func comparing(l1, l2 []int) {
	if len(l1) != len(l2) {
		fmt.Println(0)
	} else {
		c := 0
		for i := range l1 {
			if l1[i] == l2[i] {
				c++
			}
		}
		if c == len(l1) {
			fmt.Println(1)
		} else {
			fmt.Println(1)
		}

	}

}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	v, _, _ := reader.ReadLine()
	vv := strings.TrimRight(string(v), "\n")
	t_temp, _ := strconv.ParseInt(vv, 10, 64)

	for i := 0; i <= int(t_temp); i++ {

		var new = [][]int{}
		for g := 0; g < 2; g++ {
			v, _, _ := reader.ReadLine()
			vv := strings.TrimRight(string(v), "\n")
			nu, _ := strconv.ParseInt(vv, 10, 64)

			li := []int{}
			for j := 0; j < int(nu); j++ {

				v, _, _ := reader.ReadLine()
				vv := strings.TrimRight(string(v), "\n")
				df, _ := strconv.ParseInt(vv, 10, 64)
				li = append(li, int(df))
			}
			new = append(new, li)
		}
		fmt.Println(new)
		fmt.Println(new[0])
		fmt.Println(new[1])

		comparing(new[0], new[1])
	}

}

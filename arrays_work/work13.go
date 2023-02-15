package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func duplicate(sli []int) []int {

	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range sli {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	out, _, _ := reader.ReadLine()

	word := string(out)

	Split_str := strings.Split(word, " ")

	int_slice := []int{}

	for _, v := range Split_str {
		vv, _ := strconv.Atoi(v)
		int_slice = append(int_slice, vv)
	}

	fmt.Println(duplicate(int_slice))
}

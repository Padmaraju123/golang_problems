package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func comparing(l1 [][]int) bool {
	return reflect.DeepEqual(l1[0], l1[1])

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	val, _, _ := reader.ReadLine()
	b, _ := strconv.ParseInt(string(val), 10, 64)
	var a [][]int
	for i := 0; i < int(b); i++ {
		val2, _, _ := reader.ReadLine()
		str := strings.Split(string(val2), " ")
		var b []int
		for i := 0; i < len(str); i++ {
			val, _ := strconv.Atoi(str[i])
			b = append(b, val)
		}
		a = append(a, b)
	}
	fmt.Println(comparing(a))
}

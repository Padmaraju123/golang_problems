package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gett(sl []string, nn int) []int {

	out := []int{}

	for _, i := range sl {
		k,_ := strconv.Atoi(i)
		if k > nn {
			out = append(out, k)
		}
	}
	return out

}

func main() {
	user := bufio.NewReader(os.Stdin)
	word, _ := user.ReadString('\n')
	split := strings.Split(word, " ")

	var n int
	fmt.Scanln(&n)
	fmt.Println(gett(split, n))

}

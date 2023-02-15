package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ree(sll map[string]int, ss string) map[string]int {

	slii := strings.Split(ss, " ")

	fmt.Println(slii)

	for _, dd := range slii {
		_, check := sll[dd]
		if check == true {
			delete(sll,dd)
		}

	}
	return sll

}

func main() {
	var alpha = map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
		"d": 100,
		"e": 101,
		"f": 102,
		"g": 103,
		"h": 104,
	}

	user_str := bufio.NewReader(os.Stdin)
	s, _ := user_str.ReadString('\n')

	fmt.Println(ree(alpha, s))

}

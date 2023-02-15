package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	get_word := bufio.NewReader(os.Stdin)
	get_Str, _ := get_word.ReadString('\n')
	split := strings.Split(get_Str, " ")
	final := strings.Join(split, " ")
	fmt.Println(final)
	fmt.Printf("%T %v\n", final, final)
}

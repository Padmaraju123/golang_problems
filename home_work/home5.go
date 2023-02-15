package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"strconv"
	"strings"
)

func gg(word string, cash int) {

	split := strings.Split(word, " ")
	fmt.Println(split)
	nu := []int{}
	for _, k := range split {
		vv, _ := strconv.Atoi(k)
		fmt.Printf("%T %v\n", vv, vv)
		nu = append(nu, vv)
	}
	fmt.Println(nu)

	/*le := len(word)
	fmt.Println(word)

	for i := 0; i < le; i++ {
		bb := cash / word[i]
		fmt.Printf("%d Notes = %d\n", word[i], bb)
		cash = cash % word[i]
	}*/

}

func main() {
	get := bufio.NewReader(os.Stdin)
	out, _ ,_:= get.ReadLine()
	k:=string(out)

	var c int
	fmt.Scanln(&c)
	gg(k, c)

}

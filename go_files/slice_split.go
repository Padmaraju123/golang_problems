package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func conca(ww string)[]string{
	split_ww := strings.Split(ww, " ")
	fmt.Println(split_ww)
	return split_ww
}

func main() {
	var word string
	fmt.Scanln(&word)
	fmt.Printf("%T %q %s\n", word, word, word)
	fmt.Println(word)

	fmt.Println("ddfd ddfd\ndfdfsdfs")
	fmt.Println(`dfsfsfgs\ndfadsfsfsd`) //raw string so escape character not work

	fmt.Println("the title: harry poter\nthe rating: 5/5")

	fmt.Println(`
the title: harry poter
the ratings: 5/5
	`)

	//concatenation of strings
	var total_word = "i " + "am " + word
	fmt.Println(total_word)

	user_one := bufio.NewReader(os.Stdin)
	user_two,_ := user_one.ReadString('\n')

	ott := conca(user_two)

	for _,hh := range ott{
		fmt.Println(hh)
	}

}

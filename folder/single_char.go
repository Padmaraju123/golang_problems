package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check_letter(w rune){
	if unicode.IsLetter(w){
		fmt.Printf("%c is an alphabet",w)
	}else{
		fmt.Printf("%c is not a alphabet",w)
	}
}

func main() {
	read := bufio.NewReader(os.Stdin)
	number, _, _ := read.ReadRune()
	fmt.Println(number)
	check_letter(number)


}

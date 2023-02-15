package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
func check_letter(w byte) {
	if unicode.IsLetter(w) {
		fmt.Printf("%c is an alphabet", w)
	} else {
		fmt.Printf("%c is not a alphabet", w)
	}
}*/

func main() {
	get := bufio.NewReader(os.Stdin)
	out, error := get.ReadByte()
	_ = error
	fmt.Println(out)
	fmt.Println(rune(out))
	//check_letter(out)
}

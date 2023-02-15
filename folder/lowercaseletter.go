package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// to change the character is lowercase or not
func main() {
	ch := bufio.NewReader(os.Stdin)
	val, size, error := ch.ReadRune()
	fmt.Printf("the data type %T and the value is %v\n", size, size)
	fmt.Printf("the data type %T and the value is %v\n", error, error)
	fmt.Printf("the data type %T and the value is %v\n", val, val)

	if unicode.IsLower(val) {
		fmt.Println("you entered the lowercase letter")

	}else if unicode.IsUpper(val){
		fmt.Println("you entered upper case letter")
	}else{
		fmt.Println("you entered unknown value")
	}

}

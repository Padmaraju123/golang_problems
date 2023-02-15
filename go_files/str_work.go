package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b' //rune data type so int32
	fmt.Printf("Type: %T,Value: %d\n", var1, var1)
	fmt.Printf("Type: %T,Value: %d\n", var2, var2)

	//this is byte characters so we get length based on no of characters and it have ASCII values
	var word = "tara"
	var le = len(word)
	fmt.Println(le)

	for i:=0;i<le;i++{
		fmt.Printf("%c\n",word[i])
	}

	//non ASCII characters word may different length
	var word1 = "ĀĒĪŌŪȲḠ"
	fmt.Println(len(word1))
	//to find no of runes/characters for non-ascii words
	n := utf8.RuneCountInString(word1)
	fmt.Println(n)

	//slicing ascii string returns bytes not runes
	var sli_ww = "i am ok"
	fmt.Println(sli_ww[2:])

	//slicing a non-ascill string 
	fmt.Println(word1[2:])



}

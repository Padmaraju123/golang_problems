package main

import "fmt"

func main(){
	word := "This is an example."
	runes := []rune(word)
	fmt.Printf("%T %v\n",runes,runes)
	le := len(runes)

	for i:=0;i<le;i++{
		fmt.Printf("%T %c\n",runes[i],runes[i])
		fmt.Printf("%T %v\n",string(word[i]),string(word[i]))
		
	}
}
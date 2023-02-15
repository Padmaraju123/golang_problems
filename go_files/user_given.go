package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//it's print us single word after space it not 
	var give_word string 
	fmt.Scanln(&give_word)
	fmt.Println(give_word)

	//to print enter name 
	given := bufio.NewReader(os.Stdin)

	var arrr = []string{}
	for i:=0;i<3;i++{
		outt,_:=given.ReadString('\n')
		arrr=append(arrr, outt)
	}

	fmt.Println(arrr)
}
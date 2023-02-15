package main

import (
	"fmt"
	"os"
)

// The os package contains the Args array, a string array that contains all the command-line arguments passed
//slice of string array is os.args
//we need to give values in terminal by command way only "go run filename.go"
//first argument gives  program name
func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:]) //GIVES YOU SLICE OF STRING ELEMENTS WHICH YOU HAVE ENTERED 
	fmt.Println(len(os.Args)) //gives the size of the string slice

	//iterating slice of string values
	for _,val:= range os.Args[1:]{
		fmt.Printf("the data type is %T and the value is %v\n",val,val)
	}
}

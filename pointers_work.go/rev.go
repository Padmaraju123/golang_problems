package main

import "fmt"

var num int
func main(){
	//here num is variable like temporary memory to store the value and the store place has address 
	num = 100
	
	fmt.Println(&num)
}
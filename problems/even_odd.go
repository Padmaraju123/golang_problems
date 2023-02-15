package main

import "fmt"

func even_odd(num int)string{
	if num%2==0{
		return "this is even number"
	}
	return "this is odd number"
}

func main(){
	var num int 
	fmt.Scanln(&num)
	fmt.Println(even_odd(num))	
}
package main

import "fmt"

//passing the pointers to functions to change the values
//if you pass variable to function it won't change until it has return keyword
//but if you pass pointer of the varible it works

//normal function
func change1(a int)int{
	a = 200
	return a

}

//pointer pass function 
func change2(a *int){
	*a = 300
}

func main(){

	var num = 100

	fmt.Println(num)

	//pass the variable
	change1(num)
	fmt.Println(num)

	//passing the pointer of the num variable
	change2(&num)
	fmt.Println(num)

}
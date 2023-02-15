package main

import "fmt"

func f1(a ,b int)int{
	return a+b

}

func f2(a,b int)int{
	return a*b

}

func main(){

	fmt.Println("here we are using defer keyword")

	defer fmt.Printf("the sum of a,b : %d\n",f1(10,20))

	fmt.Printf("the product of a,b: %d\n",f2(1,0))

	defer fmt.Println("hello")

}
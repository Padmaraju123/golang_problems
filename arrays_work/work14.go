package main

import "fmt"

func main(){
	var arr1 [3]int
	fmt.Printf("%T %v\n",arr1,arr1)

	arr1[0] = 100
	fmt.Printf("%T %v\n",arr1,arr1)

}
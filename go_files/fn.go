package main

import "fmt"

func Jii(arry ...int){
	var le = len(arry)

	for i:=0; i<le ; i++{
		fmt.Println(arry[i])
	}
	
}

func main(){
	arry := []int{23,4,5}
	fmt.Println(arry)
	Jii(arry...)
}


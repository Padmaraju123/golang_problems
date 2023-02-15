package main

import (
	"fmt"
	"sort"
)

func main(){
	var ski = []int{100,200,300}

	for i:=0;i<5;i++{
		ski = append(ski)
	}
	fmt.Println(ski)
	ski = append(ski[1:],0)
	fmt.Println(ski)

	//creating a slice with make function

	new_sk := make([]int,3)
	new_sk[0] = 100
	new_sk[1] = 500
	new_sk[2] = 200

	fmt.Println(new_sk)
	
	//sort the elements in increasing order
	sort.Ints(new_sk)

	fmt.Println(new_sk)

	//how to remove the element of the slice by index position
	//by append method we can perform
	var index int = 1
	new_sk = append(new_sk[:index],new_sk[index+1:]...)
	fmt.Println(new_sk)
	
	
}
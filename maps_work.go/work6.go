package main

import "fmt"

func mmm(b int)map[int]int{
	var em = map[int]int{}

	for i:=1;i<=b;i++{
		em[i] = i*2
	}

	return em


}

func main(){
	var ss int 
	fmt.Scanln(&ss)
	fmt.Println(mmm(ss))

}
package main

import "fmt"

func main(){
	var s11 = []int{1,2,3}

	var s22 = make([]int,len(s11))

	s33 := copy(s22,s11)
	fmt.Println(s11,s22,s33)

}
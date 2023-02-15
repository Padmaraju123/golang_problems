package main

import "fmt"

func facc(num int)[]int{
	var out = []int{}
	for i:=1;i<=num;i++{
		if num%i==0{
			out=append(out,i)
		}

	}
	return out
}

func main(){
	var num int
	fmt.Scanln(&num)
	fmt.Println(facc(num))
}
package main

import "fmt"

func function2(s []int,el int)int{
	var out int
	for j,_ := range s{
		if s[j] == el{
			out = j
			break
		}
	}
	return out

}

func function1(val,ele int)int{
	sli := make([]int,val)
	for i:=0;i<val;i++{
		fmt.Scanln(&sli[i])
	}
	out := function2(sli,ele)
	return out
}

func main(){
	var size,element int
	fmt.Scanln(&size)
	fmt.Scanln(&element)
	fmt.Println(function1(size,element))

}
package main

import (
	"fmt"
	"strings"
)

func pattern(num int){
	j:=1

	for j<num+1{
		pat1 := strings.Repeat(" ",num-j)
		pat2 := strings.Repeat("x-",j)
		fmt.Println(pat1+pat2)
		j++
	}

	i:=1

	for i<num{
		pat3 := strings.Repeat(" ",i)
		pat4 := strings.Repeat("x-",num-i)
		fmt.Println(pat3+pat4)
		i++

	}
}

func main(){
	var num int
	fmt.Scanln(&num)
	pattern(num)
}
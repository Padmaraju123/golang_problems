package main

import (
	"fmt"
	"strings"
)

func hollow_diamond(num int){
	for i:=0;i<num;i++{
		out1 := strings.Repeat(" ",num-(i+1))
		if i==0{
			fmt.Println(out1+"x")
		}else{
			out2 := strings.Repeat(" ",i+1)
			fmt.Println(out1+"x"+out2+"x")
		}

	}
}

func main(){
	var num int
	fmt.Scanln(&num)
	hollow_diamond(num)
}
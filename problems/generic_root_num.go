package main

import "fmt"

func generic_root_num(num int){

	sum := 0
	for num>10{
		for num>0{
			re := num%10
			num = num/10
			sum+=re
		}
		num = sum
		sum = 0
	}
	fmt.Println(num)
}

func main(){
	var num int 
	fmt.Scanln(&num)
	generic_root_num(num)
}
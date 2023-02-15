package main

import "fmt"

func first_digit_num(num int)int{
	
	for num>10{
		num = num/10
	}
	return num

}

func main(){
	var num int 
	fmt.Scanln(&num)
	fmt.Println(first_digit_num(num))
}
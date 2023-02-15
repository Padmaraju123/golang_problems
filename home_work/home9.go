package main

import "fmt"

func first(n int)int{
	for n>10{
		n = n/10
	}
	return n
}

func main(){
	var num int
	fmt.Scanln(&num)
	fmt.Println(first(num))

}
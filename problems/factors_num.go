package main

import "fmt"

func factors_num(num int){

	for i:=1;i<num+1;i++{
		if num%i==0{
			fmt.Println(i)
		}
	}

}
func main(){
	var num int 
	fmt.Scanln(&num)
	factors_num(num)
}
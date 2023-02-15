package main

import "fmt"

func ssum(n int)int{
	var sumw = 0
	for i:=1;i<n+1;i++{
		sumw+=i
	}

	return sumw

}

func main(){
	var nun int
	fmt.Scanln(&nun)
	fmt.Println(ssum(nun))
}
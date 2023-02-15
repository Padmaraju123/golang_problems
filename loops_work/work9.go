package main

import "fmt"

func even_sum(n int)int {
	var ss = 0
	for i:=1;i<n+1;i++{
		if i%2==0{
			ss+=i

		}
	}
	return ss

}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(even_sum(n))
}

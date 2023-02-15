package main

import "fmt"

func repp(m,n int)[]int{
	var k = []int{m}
	i:=0
	for i<n{
		k=append(k,m)
	}
	return k
}	

func main() {
	var m,n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Println(repp(m,n))

}
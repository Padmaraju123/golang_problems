package main

import "fmt"

func divisiblebyM_N(maxi, m, n int) {
	for num:=1;num<=maxi;num++{
		if num%m == 0 && num%n == 0 {
			fmt.Printf("the %d is divisible by %d & %d\n", num, m, n)
		}
	}

}

func main() {
	var maxi, m, n int
	fmt.Scan(&maxi, &m, &n)
	divisiblebyM_N(maxi, m, n)
}

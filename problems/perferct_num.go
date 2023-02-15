package main

import "fmt"

var sum int

func perferct_num(mini,maxi int) {

	for mini<=maxi{
		sum = 0
		for i := 1; i < mini; i++ {
			if mini%i == 0 {
				sum += i
			}
		}
		if sum == mini {
			fmt.Printf("%d is perfect num\t",mini)
		}
		mini++
	}
}

func main() {
	var mini,maxi int
	fmt.Scanln(&mini,&maxi)
	perferct_num(mini,maxi)

}

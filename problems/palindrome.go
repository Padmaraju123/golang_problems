package main

import "fmt"

func palindrome(mini, maxi int) {
	var check, rev,re int
	for mini <= maxi {
		check = mini
		rev = 0
		for check > 0 {
			re = check % 10
			rev = rev*10 + re
			check = check / 10
		}
		if rev == mini {
			fmt.Println(mini)
		}
		mini++
	}

}


func main() {
	var mini, maxi int
	fmt.Scan(&mini, &maxi)
	palindrome(mini, maxi)

}

package main

import "fmt"

func print_reverse_n(mini,maxi int) {
	for maxi >= mini {
		fmt.Print(maxi,"\t")
		maxi--
	}
}

func main() {
	var mininum,maxinum int
	fmt.Scan(&mininum,&maxinum)
	print_reverse_n(mininum,maxinum)
}

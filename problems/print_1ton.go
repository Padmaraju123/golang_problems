package main

import "fmt"

func print_1toN(startnum, endnum int) {
	for startnum <= endnum {
		fmt.Print(startnum, "\t")
		startnum++
	}
}

func main() {
	var startnum, endnum int
	fmt.Scan(&startnum, &endnum)
	print_1toN(startnum, endnum)
}

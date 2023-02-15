package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	var st = []string{}
	st = os.Args[1:]
	fmt.Println(st)
	


}

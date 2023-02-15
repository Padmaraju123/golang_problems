package main

import (
	"fmt"
	"strconv"
	"strings"
)

func extract(w string) int {
	var ot int
	st := strings.Split(w, "")
	for _, j := range st {
		gh, _ := strconv.Atoi(j)
		ot += gh
	}
	return ot
}

func main() {
	var word string
	fmt.Scanln(&word)
	fmt.Println(extract(word))
}

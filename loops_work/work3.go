package main

import (
	"fmt"
	"strings"
)

func tri(nn int) {
	for i := 1; i < 5; i++ {
		fmt.Println(strings.Repeat("* ", i))
	}
}

func main() {
	var nn int
	fmt.Scanln(&nn)
	tri(nn)

}

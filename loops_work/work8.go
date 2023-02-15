package main

import (
	"fmt"
	"strings"
)

func ptt(n int) {
	i := 1
	for i < n+1 {
		fmt.Println(strings.Repeat("* ", i))
		i++
	}
}

func main(){
	var k int
	var tr int
	fmt.Scanln(&k)
	fmt.Scanln(&tr)
	for h:=0;h<k;h++{
		ptt(tr)
	}
}

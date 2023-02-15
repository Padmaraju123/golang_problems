package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func fgg(s1, s2 string) {
	s11 := strings.Split(s1, " ")
	s22 := strings.Split(s2, " ")
	sort.Strings(s11)
	for i:=0;i<len(s11);i++{
		fmt.Printf("%v %v\n",s11[i],s22[i])
	}
	

}

func main() {
	get := bufio.NewReader(os.Stdin)
	g, _ := get.ReadString('\n')
	h, _ := get.ReadString('\n')
	fgg(g, h)

}

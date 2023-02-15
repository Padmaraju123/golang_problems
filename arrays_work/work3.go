package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vv(n int,kk []string)[]string{
	var out = []string{}
	for i:=0;i<n;i++{
		var j int
		fmt.Scanln(&j)
		out = append(out,kk[j])
	}
	return out
	

}

func main() {
	var n int
	fmt.Scanln(&n)
	oo := bufio.NewReader(os.Stdin)
	df,_:= oo.ReadString('\n')
	sf := strings.Split(df," ")
	for _,gh:= range vv(n,sf){
		fmt.Println(gh)
	}
}

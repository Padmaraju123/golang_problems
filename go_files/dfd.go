package main

import (
	"fmt"
	"strings"
)

func joinStr(raju ...string) string {
	return strings.Join(raju,"#")
}

func main() {
	fmt.Println(joinStr("jjs","sdsd","Sdsd"))
	fmt.Println(joinStr("jj","sds","Sds"))
	fmt.Println(joinStr("j","sd","Sd"))

}
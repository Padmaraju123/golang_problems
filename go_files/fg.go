package main

import (
	"fmt"
	"strings"
)
//with out using variadic function for slice
func joinstr(elements []string)string {
	fmt.Println(elements)
	return strings.Join(elements,"")
}

//with using variadic function for slice
func ji(e ...string)string{
	return strings.Join(e,"#")

}

func main() {

	elements := []string{"2","2"}
	fmt.Println(joinstr(elements))

	var ww = []string{"A","sd","sdsd"}
	fmt.Println(ji(ww...))
	fmt.Println(ji("DWWE","sdSDS","SDSDSDSD"))

}

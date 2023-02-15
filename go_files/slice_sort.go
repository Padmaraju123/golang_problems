package main

import (
	 "fmt"
	"sort"
)

func main(){
	aruu := []string{"dgdf","Asa","yrtr","erer"}
	sgh := []int{23,34,5454}

	//sorting of string
	sort.Strings(aruu)
	sort.Ints(sgh)

	//gives false or true whether sorted or not
	check1 := sort.StringsAreSorted(aruu)
	check2 := sort.IntsAreSorted(sgh)

	fmt.Printf("%t %t\n",check1,check2)

	fmt.Printf("%#v\n",aruu)
	fmt.Printf("%#v\n",sgh)

	slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',  
	'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}

	fmt.Print(slice_1)

}
package main

import (
	"fmt"
)

func main(){

    var my_slice_1 = []int{0,2023,34223}

	for i:=0;i<len(my_slice_1);i++{
		fmt.Println(my_slice_1[i])
	}

	for we,tr := range my_slice_1{
		_ = we
		fmt.Println(tr)
		
	}

	for _,gt := range my_slice_1{
		fmt.Println(gt)
	}
}				   
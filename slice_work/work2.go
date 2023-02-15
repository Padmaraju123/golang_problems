package main

import (
	"fmt"
)

func main() {
	var ary [2]int //unintialised array contains zeros
	fmt.Printf("%T\n", ary)
	fmt.Println(ary)
	fmt.Println(len(ary))

	//how to compare two arrays it gives true if it unintialised but false to intialised array but accual it wont work for array it works only slices

	//we can  assign values to unintialised array
	ary[0] = 200
	fmt.Println(ary)

	//how to create a array with make function
	new_A := make([]int, 4)
	fmt.Println(new_A)

	new_A[0] = 1002
	new_A[1] = 323
	new_A[2] = 2323
	new_A[3] = 2323

	fmt.Println(new_A)

	//create a array with different way
	type names []string
	friends := names{"rs", "dfd", "sdsdsds"} //names is class and friends is object

	fmt.Println(friends)
	fmt.Printf("%T\n", friends)

	for _, i := range ary {
		fmt.Println(i)
	}

	var slii []int //unintialised slice contains empty or nil
	fmt.Printf("%T\n", slii)
	fmt.Println(slii)
	fmt.Println(len(slii))

	//how to compare two arrays it gives true if it unintialised but false to intialised array
	fmt.Println(slii == nil)

	//we can't assign values to uninitialised slice
	//slii[0] = 303
	//fmt.Println(slii)

	//create a slice with make function

	new_ss := make([]string, 3)
	//here we can put values in in with length of 3 by using index way but we can put more of it by append way
	new_ss[0] = "dsfsfds"
	new_ss[1] = "fgsfgsag"
	new_ss[2] = "sghsgs"
	//new_ss[3] = "fafdsdfs"
	fmt.Println(new_ss)

	//by looping we can put more values
	for i := 0; i < 2; i++ {
		var ggg = "yono"
		new_ss = append(new_ss, ggg)
	}

	fmt.Println(new_ss)

	//class way to create a slice
	type ok []int
	grt := ok{100, 2003, 4343}
	fmt.Println(grt)
	fmt.Printf("%T\n", grt)

	var ar = [2]string{"raju", "sdsds"} //intialised array
	rrt := [2]string{"raju"}
	fmt.Println("here compare")
	fmt.Println(ar == rrt)
	fmt.Println(ar)
	fmt.Printf("%T\n", ar)
	fmt.Printf("%#v\n", ar)
	fmt.Printf("%v\n", ar)
	fmt.Println(len(ar))

	var sj = []string{"sdsd", "dfrtdf"} //intialised slice
	fmt.Println(sj)
	fmt.Printf("%T\n", sj)

	//how to compare intialised slices with elements

	one, two := []int{1, 2, 2}, []int{1, 2, 2}

	//fmt.Println(one == two) it will case error
	//to compare it we go with looping
	var check bool = true
	for i, gh := range two {
		if gh != one[i] {
			check = false
			break
		}
	}	

	if check {
		fmt.Printf("%v and %v are equal slices\n", one, two)
	} else {
		fmt.Println("not equal")
	}

	

	//how to access the particular element throgh index value
	fmt.Println(sj[0])

	//access the elements throgh loop
	for _, sss := range sj {
		fmt.Println(sss)
	}

	//how to update the values by indexing
	sj[0] = "omg"
	fmt.Println(sj)

	//how to assign slice to another slice
	var get_gh []string
	get_gh = sj
	fmt.Println(get_gh)

	//if you change something in either one it will effect to other one
	get_gh[0] = "changed"
	fmt.Println(get_gh)
	fmt.Println(sj)

	//how to copy one slice to another slice we dont have copy method for looping only we can
	var cp_s []string
	for _, sdds := range sj {
		cp_s = append(cp_s, sdds)
	}

	fmt.Println(cp_s)

	//if you change either one it won't affect another
	cp_s[0] = "unchanged"
	fmt.Println(cp_s)
	fmt.Println(sj)

	//how to append values to slice
	numbers := []int{2,3}
	var new_p = []int{}

	//append the values to different slice
	new_p = append(new_p, 200)

	//append the values to same slice
	numbers = append(numbers, 200)

	fmt.Println(numbers)
	fmt.Println(new_p)

	//append a lot values once

	new_p = append(new_p,2323,34343,34343)
	fmt.Println(new_p)

	//append a whole slice to another slice

	numbers = append(numbers , new_p...)     //here  we need to like this because append is variadic function
	fmt.Println(numbers)

	//copy a slice to another slice with copy function
	src := []int{100,200,300,400}
	drc := make([]int ,len(src))
	gdf := make([]int ,2)
	kk := copy(gdf,src)
	nn := copy(drc,src)  //it gives no of elements copied

	fmt.Println(src,drc,gdf,nn,kk)

	//slicing a slice within same slice
	var first = []int{23,5,6,77}
	fmt.Println(first)
	first = first[1:]
	fmt.Println(first)
	//slicing a slice with different slice

	second := first[1:]
	fmt.Println(second)

}

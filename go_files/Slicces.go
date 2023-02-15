package main

import "fmt"

func main(){ 
	var l []int  //unintialised slice 

	fmt.Println(len(l))
 
	var sl = []int{2,3}      // initialise slice

	fmt.Println(len(sl))

	//l[0] = 10000    // for nil slice we can't add elements

 
	fmt.Printf("the list has look like this %v\n",sl)

	fmt.Printf("the list has look like this %#v\n",sl) //it gives us a list with nil elements but in arrays it gives us zeros in list

	if sl == nil{
		fmt.Println("NOOO")
	}else{
		fmt.Println("YESS")
	}

	//we can make a slice with make built-in function
	slii := make([]int,3) // 3 is length 

	fmt.Println(slii)

	//create a slice 

	type items []string

	gets := items{"pepsi","maaza"}

	fmt.Println(gets)

	//accessing the elements of slices

	fmt.Println(gets[1])

	//updating the elements of slice

	gets[0] = "Spirit"
	fmt.Println(gets)

	for ind , value := range gets{
		fmt.Println(ind,value)
	}

	//how to assign a slice into empty slice

	var emp []string
	emp = gets
	fmt.Println(emp)
}
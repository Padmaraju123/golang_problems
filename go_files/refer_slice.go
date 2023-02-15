package main

import "fmt"

func main(){

	/*Modifying Slice: As we already know that slice is a reference type it can refer an underlying array.
	 So if we change some elements in the slice, then the changes should also take place in the referenced array.
	 Or in other words, if you made any changes in the slice, 
	 then it will also reflect in the array as shown in the below example:*/

	var ary1 = [4]int{10,343,348,34}
	var sli1 = ary1[1:]
	fmt.Printf("%#v\n",ary1)
	fmt.Printf("%#v\n",sli1)

	sli1[1] = 1000
	sli1[2] = 3000

	fmt.Printf("%#v\n",ary1)
	fmt.Printf("%#v\n",sli1)

	ary1[1]= 10000
	
	fmt.Printf("%#v\n",ary1)
	fmt.Printf("%#v\n",sli1)

   
}
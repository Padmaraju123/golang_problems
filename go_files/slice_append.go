package main

import "fmt"

func main(){
	//append the element at the end of the slice
	var sll = []int{2,4,5}

	fmt.Printf("%#v\n",sll)

	sll = append(sll,1000)  //here i append the value to same slice

	new_sll := append(sll,2000)   //append into new_Slice

	fmt.Printf("%#v\n",sll)
	fmt.Printf("%#v\n",new_sll)

	//appending more elements 
	sll = append(sll,300,400,100)

	fmt.Printf("%#v\n",sll)

	//how to append one slice elements into another slice 
	old_Sl := []int{233,323}
	sll = append(sll, old_Sl...)     //here we need to append same data types
	fmt.Printf("%#v\n",sll)
}
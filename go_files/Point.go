package main

import "fmt"

func main() {
	//the pointer gives the address of the variable through this address we can access the data stored in variable
	name := "raju"
	//here it prints the address the address is hexadecimal 
	fmt.Println(&name)
	//create a pointer 
	var num int = 100
	ptr := &num
	fmt.Printf("the type of pointer %T and the value of pointer %v\n",ptr,ptr)
	//to print the address of the num
	fmt.Printf("the address of num %p\n",ptr)
	//the pointer has address it self
	fmt.Printf("the adress of ptr is %p\n",&ptr)

	//declare a pointer with zero values/unintialised 
	var ptr1 *float64

	//now i declare a value to some variable do something
	//var cc = 100
	//ptr1 = &cc //it will case error because the pointer ptr1 gives the varibles which has float values only
	//fmt.Println(ptr)

	//now same thing but correct value
	var ss = 1212.23232
	ptr1 = &ss
	fmt.Println(ptr1)

	//another way to create a pointer by new function
	ptr2 := new(string)

	//now check1
	var gd = "rajjj"
	ptr2 = &gd
	fmt.Println(ptr2)

	//changing the value of variable by pointer because it is address
	fmt.Println(*ptr2)           //* operator we can value from the variable by address
	*ptr2 = "padmaraju"          // gd == "padmaraju"
	fmt.Println(gd,*ptr2)      //* is dereferencing operator

	//pointer to pointer
	var s = 232
	p1 := &s // pointer1
	p2 := &p1
	fmt.Printf("the value of p1: %v, address of the p1: %v\n",p1,p2)
	fmt.Printf("the value of p2: %v, address of the p2: %v\n",p2,&p2)
	
	//dereference the pointer to pointer
	//i can access by value of the s 
	fmt.Println(**p2)
	**p2 = 2322       //now i am changing the value of s indirectly by address
	fmt.Println(s)

	//increment the s value
	**p2++
	fmt.Println(s)

	//comparing the pointers
	var poi *int

	aa := 100
	poi = &aa

	bb := 200 
	poy := &bb 

	//if the two pointers are used for same variable then it returns true else false even though the variables values are same
	fmt.Println(poi == poy)

	pok := &aa 
	fmt.Println(poi == pok)




}

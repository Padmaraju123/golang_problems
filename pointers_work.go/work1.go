package main

import "fmt"

func main() {
	name := "taju"

	//if i want the address of the name in memoery
	fmt.Printf("\n%T %v\n\n",&name,&name)



	//how to declare and intialse pointers
	var x int = 2
	ptr := &x        //*int = pointer to int
	fmt.Printf("the data type %T and the value is %p\n\n", ptr, ptr)
	fmt.Printf("the address of x is %p\n\n", ptr)

	//now the ptr address variable has address
	fmt.Printf("the address of ptr is %p\n\n", &ptr)

	//to unintialise the pointer with nil value
	var ptr1 *float64
	fmt.Printf("the value is %T and the address is %p %v\n\n", ptr1, &ptr1,ptr1)

	
	//here we use the ptr1 to get the address of values float64 only
	/*var float_vlaue float64 = 1034.4343
	ptr1 = &float_vlaue
	fmt.Println(ptr1)
	fmt.Println(&ptr1)

	//another way to fix pointer to datatype by using new function
	ptr2 := new(string)

	string_value := "padmaraju"

	ptr2 = &string_value
	fmt.Printf("the data type of ptr2 %T and the address of ptr2 %p\n",ptr2,ptr2)

	//to change the variable value by address this is dereferencing operator
	*ptr2 = "siddanatham padmaraju"      //string_value = "siddanatham padmaraju"
	fmt.Println(string_value,*ptr2)*/

	

}

package main

import "fmt"

func main(){
	//pointer syntax

	var address_num *int       //here ptr used to store the address of integers values only
	fmt.Printf("%T %v\n",address_num,address_num)   //it returns data type as *int and value is nil


	//now assign value to variable
	var num = 100
	address_num = &num
	fmt.Printf("%T %v\n",address_num,address_num)   //it return data type as *int and value is hexadecimal number


	//Dereferencing a pointer means getting the value inside the address the pointer holds. 
	//If we have a memory address, we can dereference the pointer to that memory address to get the value inside it. 
	//Here is the same example showing the dereference operation using the star(*) operator.

	value := *address_num
	fmt.Printf("%T %v\n",value,value)

	//A pointer variable can store even a pointers address since a pointer is also a variable just like others.
	address_of_address := &address_num
	fmt.Printf("%T %v\n",address_of_address,address_of_address)  //it returns data type as **int and value is hexadecimal

}

package main

import "fmt"

func main(){
	var a uint8 = 10    //uint8 has range 0 to 255  and uint8 is underlying type
	fmt.Printf("%T %v\n",a,a)

	//byte is alias of uint8
	var b byte = 200
	fmt.Printf("%T %v\n",b,b)

	//we can perform operations with same data type 
	fmt.Printf("%T %v\n",a+b,a+b)

	//we can rename with our name to the defined data types nothing but aliasing
	type second = uint8

	var num second = 255
	fmt.Printf("%T %v\n",num,num)

	type num1 int8     //here int8 is the underlying type of the num1
	type num2 num1     // here also int8 is the underlying type of the num2

	var value num1 = 23         // here num1 is the data type of value not int8

	var value1 int8 = 20
	//out := value1+value    //not preform because different data type for value it is num1 and for value1 it is int8
	//if you do convert it into int8
	out := value1+int8(value)
	fmt.Println(out)

	//out1 := value+value1  // here also error because different data types
	out1 := value+num1(value1)     //convert to num1 
	fmt.Println(out1)

}
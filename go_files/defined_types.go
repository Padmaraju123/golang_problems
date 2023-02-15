package main

import "fmt"

type km float64
type mile float64

func main(){
	//type age int     int is the underline type
	//type oldage age    int is the underline type 
	//type veryoldage age  int is the underline type

	type speed uint         //type is keyword and speed is identifier and uint is type of speed

	var s1 speed  = 23                 //here var is keyword and s1 is identifier and speed is type of s1 that is uint
	var s2 speed  = 10 

	fmt.Println(s1-s2)

	var s3 uint = 24
	//s2 = s3  it occurs error because s3 and s1  are different types
	//to convert and do operation
 
	s3 = uint(s1)
	fmt.Println(s3)

	var s4 speed = speed(s3)

	fmt.Println(s4)


	var AtoB km = 100
	var miles mile 
	//miles = AtoB/0.6   //here it occurs error that to it has same type float64 so convert accordingly
	miles = mile(AtoB)/0.6 

	fmt.Println(miles)
}
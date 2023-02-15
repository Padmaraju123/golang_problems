package main

import "fmt"

func main(){
	//pointer address is pointer to pointer
	a := 100
	p1 := &a
	p2 := &p1

	fmt.Printf("the address of a %p and address of p1 %p\n",p1,p2)

	//the dereference to pointer to pointer

	*p1 = 200
	//it will change the value of  a = 200 because *p1 is nothing but a
	fmt.Println(a)
	
	**p2 = 300 
	//it will change the value of a = 300 because **p2 is nothing but a
	fmt.Println(a)
   
	//compare the pointers
	//th comparsion possible only if the pointers are point to same variable 
	var point1 *int 

	y := 128
	point1 = &y

	z := 12
	point2 := &z

	point3 := &y

	fmt.Println(point1==point2)
	fmt.Println(point1==point3)

	

}
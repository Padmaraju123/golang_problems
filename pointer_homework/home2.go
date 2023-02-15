package main

import "fmt"

func main(){
	//address of the address or pointer of the pointer
	var p1 *int
	a := 100
	p1 = &a
	p2 := &p1
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(a)

	//changing the value of a through pointer
	*p1 = 200
	fmt.Println(a)

	//change the value of a through pointer of pointer
	**p2 = 300
	fmt.Println(a)

	//comparsion of the pointers only possible it has shared a common variable
	var check = "ramesh"

	point1 := &check

	point2 := &check

	fmt.Println(point1==point2)
}
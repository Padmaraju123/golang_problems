package main

import "fmt"

//passing a struct to function

func get1(b material) {
	b.mobile = "realme"
	b.cost = 6000
}

func get2(c *material){
	(*c).mobile = "vivo"
	(*c).cost = 3453
}

//create a  struct
type material struct {
	mobile string
	cost   int
}

func main() {
	obj := material{
		mobile: "samsung",
		cost:   5000,
	}

	get1(obj)
	fmt.Println(obj)
	//it wont change the values

	get2(&obj)
	fmt.Println(obj)




}

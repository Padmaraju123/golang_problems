package main

import "fmt"

func main(){
	//anonymous struct does not have type for this we create multiple objects to all

	diana := struct{
		firstname,lastname string
		age int
	}{
		firstname:"padmarju",
		lastname: "Siddanatham",
	}

	//to access the values in anonymous struct
	fmt.Println(diana.firstname)
	fmt.Println(diana.age)

	//to create multiples objects just like normal structs by using anonymous struct

	type Book struct{
		string
		float64
		int
	}

	address1 := Book{"ramesh",23.4,100}
	//access the values
	fmt.Println(address1.float64)

	//how to change the values
	address1.string = "raki"
	fmt.Println(address1.string)

	//we can mixup above structs and normal structs

	type circus struct{
		amount int
		name string
		bool
	}

	day := circus{500,"indra",true}
	fmt.Println(day.bool)
}
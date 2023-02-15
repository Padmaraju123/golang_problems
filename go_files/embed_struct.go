package main

import "fmt"

func main(){

	type contact struct{
		email,address string
		phone int
	}

	type employee struct{
		name string
		salary int
		contactInfo contact        //here we assign type as struct
	}

	john := employee{

		name:"John keller",

		salary: 4000,

		contactInfo: contact{
			email:"padsd@gmail.com",
			address: "asdsd sd",
			phone:232323232,
		},

	}

	fmt.Printf("%+v\n",john)
	//how to access the embedded variable
	fmt.Println(john.contactInfo.phone)
	//Update the values of embedded variabel
	john.contactInfo.phone = 12345
	fmt.Println(john.contactInfo.phone)


}
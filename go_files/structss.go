package main 

import "fmt"

func main(){
	 
	//declaring a struct or structure
	type my_Stt1 struct{
		name string
		age int
		gender string
		points float32
	}

	type my_Stt2 struct{
		name, gender  string   //we can do like this if the identifiers have same data type
		age int
		points float32
	}

	//it is way to assign values to the particular variables of the struct but it is not a good way to do
	my1 := my_Stt1{"raju",24,"male",23.5}
	fmt.Println(my1)
	fmt.Printf("%T\n",my1)

	//best way to give values to the identifiers

	my2 := my_Stt1{
		name:"padmraju",
		age : 25,
		gender : "male",
		points : 233.4,
	}

	fmt.Printf("%T, %v\n",my2,my2)

	mm := my_Stt1{
		name:"rere",
	}

	fmt.Printf("%v\n",mm)

	//to update the values 
	mm.age = 100
	mm.points = 234.4
	mm.gender = "male" 

	fmt.Printf("%+v\n",mm)

	//how to get values of particular field name
	fmt.Println(my1.name)

	//we can compare the structs with double equal

	dope := my_Stt1{
		name:"padmraju",
		age : 25,
		gender : "male",
		points : 233.4,
	}

	fmt.Println(dope == my2 )

	//how to copy structs

	copy_St := dope
	//if you change the copied one it won't change the original struct
	copy_St.name = "indra"
	fmt.Println(dope,copy_St)
	//if you change the original one it wont change the copied ones
	dope.age = 2000
	fmt.Println(dope,copy_St)

	//anonymous struct 

	diana := struct{
		firstname , lastname string
		marks int
	}{
		firstname:"padmaraju",
		lastname: "siddanatham",
		marks:195,
	}

	fmt.Printf("%T\n",diana)
	fmt.Printf("%v\n",diana)
	fmt.Printf("%+v\n",diana)

	//creating structs without variable names

	type extra struct{
		string
		int
		float32

	}

	_get := extra{"super",2055,33.5}
	//to get the value
	fmt.Println(_get.string)
	//update the value
	_get.string = "andhra"
	fmt.Println(_get.string)

	type extra1 struct{
		_name string
		_age int
		bool
	}

	_gee := extra1{"chitra",34,true}
	fmt.Println(_gee)

	
}
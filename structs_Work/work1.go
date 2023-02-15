package main

import "fmt"

func main() {

	//create a struct
	type biodata struct {
		name, gender string
		age          int
	}

	//assign a values to the fields for that create object to class of biodata
	one1 := biodata{"padmaraju", "male", 34}
	//access the data just like oops
	fmt.Println(one1.name)
	

	//the above method is not a good way to represent the data
	two1 := biodata{
		name:   "padmaraju",
		gender: "male",
		age:    25,
	}

	//if you omitted some fields it takes default values for int,float is 0 for string is space

	three1 := biodata{
		name: "siddanatham",
	}
	fmt.Println(three1.gender)
	fmt.Println(three1.age)

	//now access the values
	fmt.Println(two1.gender)

	//update the values
	fmt.Println(two1.age)
	two1.age = 50
	fmt.Println(two1.age)

	//we can compare structs values by double equal sign if the fields are same and have same values
	check1 := biodata{
		name:   "padmaraju",
		gender: "male",
		age:    50,
	}

	fmt.Println(two1 == check1)

	//how to copy a struct

	copied := check1

	fmt.Println(copied)
	//if you changed copied one's it does not effect to original one because different address in memory and vice-versa

	copied.age = 100
	fmt.Println(copied.age)
	fmt.Println(check1.age)
	check1.name = "ravi"
	fmt.Println(check1.name)
	fmt.Println(copied.name)

}

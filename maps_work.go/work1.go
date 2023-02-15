package main

import "fmt"

func main(){
	//create a map data type unintialised
	var mm  map[int]string
	fmt.Println(mm)
	fmt.Println(len(mm))
	fmt.Printf("%T %#v\n",mm,mm)

	//append values to uninitialised map it is not possible because not initialised
	//mm[2] = "raju"
	//fmt.Println(mm)

	//create a map data type initialised with zero key-value pairs
	var mk = map[string]string{}
	fmt.Println(mk)
	fmt.Printf("%T %#v\n",mk,mk)

	//append values to the initialised map
	mk["name"] = "padmaraju"

	fmt.Println(mk)
	fmt.Printf("%#v\n",mk)

	//access the map values by key names
	fmt.Println(mk["name"])
	fmt.Printf("%T\n",mk["name"])

	//create a map with default key-value pairs
	new := map[string]float64{"raju":2.3,"padmaraju":34.5,"sna":56.5}
	fmt.Println(new)
	fmt.Printf("%v\n",new)
	fmt.Printf("%T\n",new)
	fmt.Printf("%#v\n",new)
}
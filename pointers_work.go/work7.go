//use pointers in function arguments

package main

import "fmt"

type class struct{
	element int
	name string
}

func get(address *int){
	//to print the value in address use *
	fmt.Println(*address)
}

func final(add_obj *class){
	add_element := &add_obj.element
	fmt.Println(*add_element)
	add_name := &add_obj.name
	fmt.Println(*add_name)

}
//above func we can modify it like this
func final2(add_obj *class){
	fmt.Println((*add_obj).element)
	fmt.Println((*add_obj).name)
	//changing the value 
	(*add_obj).name = "padmaraju"
	fmt.Println((*add_obj).name)
}

//we can return address of the variable which was created inside the function

func ok() *int{
	v := 1000
	return &v
}


func main(){
	num := 100
	address := &num
	get(address)

	obj:=class{
		element:200,
		name:"raju",
	}

	add_obj := &obj
	final(add_obj)
	final2(add_obj)

	vv := ok()
	fmt.Println(vv)
	fmt.Println(*vv)

}
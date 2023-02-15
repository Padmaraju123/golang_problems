package main

import "fmt"

//passing slice to functions

func method1(gt1 []int){
	for i := range gt1{
		gt1[i]++
	}

}


//create a slice 
type collection []int

func main(){
	get := collection{100,200,300,400,400}

	//normal function to calling
	method1(get)
	fmt.Println(get) //for slice also

	
}
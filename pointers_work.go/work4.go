package main

import "fmt"


func changeProduct(obj product){
	obj.price = 23443.45
	obj.productname = "micromax"

}

func changeProductByPointers(obj *product){
	(*obj).price = 23443.45
	(*obj).productname = "micromax"

}

//create a struct

type product struct{
	price float64
	productname string
}

func main(){

	 //now i passing struct 
	 obj := product{
		price:2334.34,
		productname: "samsung",
	 }
	 fmt.Println(obj)

	 changeProduct(obj)
	 //it won't changed actually because we are not passing pointers
	 fmt.Println(obj)

	 //here i passing pointers of struct 
	 changeProductByPointers(&obj)
	 fmt.Println(obj.price)
	 fmt.Println(obj.productname)



}
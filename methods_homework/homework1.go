package main

import "fmt"

type car struct{
	brand string
	price int
}

//normal function to do

func changecar(c car,newbrand string,newprice int){
	fmt.Println(c.brand)
	fmt.Println(c.price)
	//changing the element values of struct data type
	c.brand = newbrand
	c.price = newprice
	fmt.Println(c.brand)
	fmt.Println(c.price)

}

//receiver function
func (c *car)changecar1(newbrand string ,newprice int){
	fmt.Println((*c).brand)
	fmt.Println((*c).price)
	//changing the element values of struct data type
	(*c).brand = newbrand
	(*c).price = newprice
	fmt.Println((*c).brand)
	fmt.Println((*c).price)
	
}

func main(){
	c:=car{
		brand:"kia",
		price:40000,

	}
	changecar(c,"tata",500000)
	c.changecar1("suzuki",600000)
}
package main

import "fmt"

//normal function
func changeCar(c car,newbrand string,newprice int){
	c.brand = newbrand
	c.price = newprice

}

//receiver function
func (g car)changeCar1(newbrand string,newprice int){
	g.brand = newbrand
	g.price = newprice

}

//pointer receiver function

func (f *car)changeCar2(newbrand string,newprice int){
	(*f).brand = newbrand
	(*f).price = newprice
}

type car struct{
	brand string
	price int
}

func main(){
	obj := car{
		brand : "tata",
		price : 5000000,
	}
	changeCar(obj,"renault",6000000)
	fmt.Println(obj)
	//now there is no change for this 

	obj.changeCar1("opel",70000000)
	fmt.Println(obj)
	//now also no change it 

	(&obj).changeCar2("tesla",8000000)
	fmt.Println(obj)
	//now its changed

	//we can also do this 
	var yourcar *car
	yourcar = &obj
	fmt.Println(*yourcar)

	//by using yourcar we can update the values because obj variable address is yourcar
	yourcar.changeCar2("VW",3434334)
	fmt.Println(*yourcar)

	(*yourcar).changeCar2("bmw",3234532434334)
	fmt.Println(*yourcar)
	
}
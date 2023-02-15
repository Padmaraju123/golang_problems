package main

import (
	"fmt"
	"math"
)

//create an interface 
type shape interface{
	//mention the methods which is used for different objects
	area() float64
	peri() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// receiver function to calculate area of the circle
func (obj1 circle) area() float64 {
	return math.Pi * math.Pow(obj1.radius,2)
}

//reciver function to calculate perimeter of the circle
func (obj1 circle) peri()float64{
	return 2*math.Pi*obj1.radius
}

//receiver function to calculate the area of the rectangle
func (obj2 rectangle) area()float64{
	return obj2.height*obj2.width

}

//receiver function to calculate the perimeter of the rectangle
func (obj2 rectangle) peri() float64{
	return 2*(obj2.height+obj2.width)
}


//using interface by using function
func print(s shape){
	fmt.Printf("Shape %#v\n",s)
	fmt.Printf("Area %v\n",s.area())
	fmt.Printf("perimeter %v\n",s.peri())
}


func main() {
	obj1 := circle{
		radius: 23.4,
	}

	print(obj1)

	obj2 := rectangle{
		width:2.5,
		height:2.5,
	}

	print(obj2)

}

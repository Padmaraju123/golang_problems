package main

import (
	"fmt"
	"math"
)

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

//method to calculate circle things
func printcircle(obj1 circle){
	fmt.Printf("area of the circle is %v\n",obj1.area())
	fmt.Printf("perimeter of the circle is %v\n",obj1.peri())
}

//method to calculate the rectangle things
func printrect(obj2 rectangle){
	fmt.Printf("area of the rectangle is %v\n",obj2.area())
	fmt.Printf("perimeter of the rectangle is %v\n",obj2.peri())
}

func main() {
	obj1 := circle{
		radius: 23.4,
	}

	printcircle(obj1)

	obj2 := rectangle{
		width:2.5,
		height:2.5,
	}

	printrect(obj2)

}

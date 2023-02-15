package main

import (
	"fmt"
	"strings"
)

//variadic function declaration
func var_f(a ...int){
	fmt.Printf("%T\n",a)
	fmt.Printf("%v\n",a)
	for _,v := range a{
		fmt.Println(v)
	}
		
}

func varrr_f(b ...int){
	b[0] = 10000    //update the value of particular index of b slice
	fmt.Println(b)
}


func SumAndProduct(c ...float64)(float64,float64){
	sum := 0.
	product := 1.

	for _,y := range c{
		sum += y
		product *= y
	}

	return sum,product
}

func Convers(age int,names ...string)string{
	fullname := strings.Join(names," ")
	fmt.Printf("%T\n",fullname)
	fmt.Println(fullname)
	retString := fmt.Sprintf("Age: %d  Full name: %s\n",age,fullname)
	fmt.Printf("%T\n",retString)
	return retString
}


func main(){
	var_f(1,2,3,4)
	var_f() 
	rj,dj := "allu","mallu"

	fmt.Printf("%s %s\n",rj,dj)

	slicc := []int{23,453}
	slicc = append(slicc, 200,304,343)    //append is in built variadic fucntion

	fmt.Println(slicc)

	var_f(slicc...)

	varrr_f(slicc...)

	fmt.Println(slicc)

	s1,p1 := SumAndProduct(2.0,3.3,1.4)
	fmt.Println(s1)
	fmt.Println(p1)

	info := Convers(25,"RAJU","KING")
	fmt.Println(info)

}
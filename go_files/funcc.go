package main

import (
	"fmt"
	
)

func Sum_1(a,b int)int{
	var s = a+b
	_ = s
	return a

}

func Allone(i,j int,x,y,z string,k,l float32)(float32,int,string){
	return k+l,i+j,x+y+z
	
}

func main(){
	fmt.Println(Sum_1(20,30))
	fmt.Println(Allone(300,200,"RAJU","SIDDAN","OK",23.4,34.5))
	out1,out2,out3 := Allone(300,200,"RAJU","SIDDAN","OK",23.4,34.5)

	fmt.Printf("%T\n",out1)
	fmt.Printf("%T\n",out2)
	fmt.Printf("%T\n",out3)

}

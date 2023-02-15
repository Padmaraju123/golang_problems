package main

import (
	"fmt"
	"time"
)

//create a receiver method
func (n names) getting(){

	for i,each_name := range n{
		fmt.Println(i,each_name)
	} 

}

//create a slice 
type names []string

func main(){
	//time.Hour is time.duration data type

	fmt.Printf("%T",time.Hour)

	const day = 24*time.Hour

	//if you want how many seconds has 24 hrs 
	seconds := day.Seconds()  //(24*time.Hour).seconds()
	fmt.Printf("the data type %T,and the no of seconds %v\n",seconds,seconds)

	//if you want how many minutes in 24 hrs 
	mins := day.Minutes()
	fmt.Printf("the data type %T, and the minutes are %v\n",mins,mins)

	sli1 := names{"raju","ravi","ramesh","rajesh"}
	//calling the receiver method
	sli1.getting()
}

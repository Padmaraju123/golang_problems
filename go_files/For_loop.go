package main

import "fmt"

//simple for loop
func main(){

	//simple for loop
	for i := 1; i < 10 ; i+=2{
		fmt.Println(i)
	}

	//infinite loop
	/*for {
		fmt.Println("Loop")
	}*/

	//for loop as while loop

	j := 0
	for j < 5 {
		fmt.Println("OK")
		j++   //it will increment by 1
	}
}


package main

import "fmt"

func leap_or_not(y int)string{
	if y%400==0 || (y%4==0 && y%100!=0) {
		return "this is leap year"
	}
	return "this is not a leap year"
}

func main(){
	var year int
	fmt.Scanln(&year)
	fmt.Println(leap_or_not(year))
}
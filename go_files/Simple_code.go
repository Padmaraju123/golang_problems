package main

import "fmt"

// var can be used outside as well as inside the function
//but := declaration is used inside the function

var a int
var b string
var c bool
var d float32

func main(){
	a = 100
	fmt.Println(a)
	b = "raju"
	fmt.Println(b)
	c = false
	fmt.Println(c)
	d = 34.5
	fmt.Println(d)

	var raju , sandhya  int  = 232,23   //for this we need give same type of values

	_ = sandhya
	fmt.Println(raju)

	one , two  := "sds",2322         // for this we can assign different data type values
	
	fmt.Println(one,two)

	var just,done = "hbdsd" , false
	fmt.Println(just,done)

	var (
		dsfd int 
		fsf  int
		ffsf int
	)

	fmt.Println(dsfd,ffsf,fsf)

	var (
		dsf int = 100 
		fs  int
		ffs int
	)

	fmt.Println(dsf,ffs,fs)

	var (
		ds int = 100 
		f  string = "rajjj"
		ff bool = false
	)

	fmt.Println(ds,ff,f)

	const csk int = 23232
	// csk = 2323232 we can't update the these identifiers because we defined as fixed value by const keyword
	fmt.Println(csk)

	fmt.Print("hello\n")
	fmt.Print("CHELO\n")
	fmt.Printf("CHELLLO\n")
	fmt.Printf("sdrddw\n")

	fmt.Println("dfsfsf")
	fmt.Println("SDSDSD")
}
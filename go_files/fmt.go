// fmt is package of Go language
//by import this package we println function/method of fmt package

package main
import "fmt"
func main(){
	name := "raju"
	fmt.Println("My name is",name)
	a,b := 10,20
	fmt.Println("SUM:",a+b,"Average:",(a+b)/2)
	//printf function from fmt package is used to give output in the console
	//but few modifications have applied when it compares to println
	fmt.Printf("my age is %d\n",25) //%d is used for integers
	fmt.Printf("my score in tenth %f\n",232.3)  //%f is used for float numbers
	fmt.Printf("my score in tenth %.3f\n",232.3)  //%.xf is used for float numbers with x digits after decimal point
	fmt.Printf("my less score is -%d\n",23)   //to give sign to the numbers
	fmt.Printf("my name is %s\n",name)         //%s is used for strings
	fmt.Printf("my Strength is %q\n","be myself")  //%q is used for quoted strings
	fmt.Printf("for all %v %v\n","cases",100)   //%v is used for all data types
	fmt.Printf("it gives us which data type %T\n",name) //%T is used gives the data type
	fmt.Printf("This statement is %t\n",2==3)    //%t is used for boolean values of the expression
	fmt.Printf("This is binary value of 5 %b\n",5)   //%b is used for binary of the numbers
	fmt.Printf("This is 8 digits binary values %08b\n",5)  //it gives the 8 digits binary value
	fmt.Printf("convert decimal to hexadecimal %x",100)  //%x is used for conversion from decimal to hexa decimal
}

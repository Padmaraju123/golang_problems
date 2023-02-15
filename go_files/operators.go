package main

import "fmt"

/* different types of operators
1)arithmetic operators
2)relational operators
3)logical operators
4)bitwise operators
5)assignment operators
6)misc operators*/

func main(){
	//arithmetic operators
	p := 10
	q := 20


	//addition 
	fmt.Printf("tHE ADDITION VALUE IS: %d\n",p+q)

	//subtraction
	fmt.Printf("the subtraction value is %d\n",p-q)

	//multiplication
	fmt.Printf("the multiplication value is %d\n",p*q)

	//division
	fmt.Printf("the division value is %d\n", p/q)

	//modulus
	fmt.Printf("the modulus value is %d\n",p%q)

	//relational operators

	//equal to (==)
	fmt.Printf("check1 %t\n",p==q)

	//not equal to(!=)
	fmt.Printf("check2 %t\n",p!=q)

	//less than (<)
	fmt.Printf("check3 %t\n",p<q)

	//greater than (>)
	fmt.Printf("check4 %t\n",p>q)

	//greater than equal to (>=)
	fmt.Printf("check5 %t\n",p>=q)

	//less than equal to (<=)
	fmt.Printf("check6 %t\n",p<=q)


	//logical operators

	//logical AND
	fmt.Printf("test1 %t\n",p!=q && p<=q)

	//logical OR
	fmt.Printf("test2 %t\n",p!=q || p<=q)

	//logical NOT
	fmt.Printf("test3 %t\n",!(p==q))





}
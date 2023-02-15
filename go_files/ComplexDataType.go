package main
import "fmt"

//Go is a statically typed programming language and this means it does type checking at compile time just like c,c++,java 

func main(){
	var one complex128 = complex(2,3)
	fmt.Println(one)

	var two complex64 = complex(45,4)
	fmt.Println(two)
	fmt.Printf("type of one identifer is %T"+" and two identifier is %T\n" ,one, two)

	var gem = 123
	fmt.Println(gem)

	var dfd int
	fmt.Println(dfd)

	var kk bool
	fmt.Println(kk)

	var df string
	fmt.Println(df)

	var look ,flower , mango  int = 12,23,534  // in case mention type then assign that values only
	fmt.Println(look,flower,mango)

	var sds ,dsd, sss = 1212,"Sdsd",false // in case not mention type then we can give different data types
	fmt.Println(sds,dsd,sss)

}
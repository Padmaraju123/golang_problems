package main
import "fmt"

func main(){
	//numeric types

	//signed integers int8,int16,int32,int64 
	//int is alias of int32,int64
	//rune is alias of int32

	//unsigned integers uint8,uint16,uint32,uint64 are for only positive integers
	//uint is an alias of uint32,uint64
	//byte is alias of uint8

	//float32,float64 are used for decimal numbers
	
	//complex64,complex128


	//the range of int8 is -128 to 127
	var n int8 = -128 
	fmt.Printf("%T %v\n",n,n)
	
	//the range of int16 is -32768 to 32767
	var n1 int16 = 32765
	fmt.Printf("%T %v\n",n1,n1)

	//the range of int32 is -2147483648 to 2147483647
	var n2 int32 = 243535343
	fmt.Printf("%T %v\n",n2,n2)

	//the range of rune is -2147483648 to 2147483647
	var p rune = 243535343
	fmt.Printf("%T %v\n",p,p)

	//the range of int64 is -9223372036854775808 to 9223372036854775807
	var n3 int64 = -3276535535353453453
	fmt.Printf("%T %v\n",n3,n3)

	//the range of int is -9223372036854775808 to 9223372036854775807 because it is alias of int32 and int64
	var n4 int = -9223372036854775808
	fmt.Printf("%T %v\n",n4,n4)



	//the range of uint8 is 0 to 255
	var n5 uint8 = 254
	fmt.Printf("%T %v\n",n5,n5)

	//the range of byte is 0 to 255
	var q byte = 254
	fmt.Printf("%T %v\n",q,q)

	//the range of uint16 is 0 to 65535
	var n6 uint16 = 65530
	fmt.Printf("%T %v\n",n6,n6)

	//the range of uint32 is 0 to 4294967295
	var n7 uint32 = 3453535345
	fmt.Printf("%T %v\n",n7,n7)

	//the range of uint64 is 0 to 18446744073709551615
	var n8 uint64 = 2343453245324532454
	fmt.Printf("%T %v\n",n8,n8)

	//the range of uint is 0 to 18446744073709551615
	var n9 uint = 3423424242524542534
	fmt.Printf("%T %v\n",n9,n9)

	//in float32 after decimal point it gives only 5 digits
	var n10 float32 = 232.343445435443544343434
	fmt.Printf("%T %v\n",n10,n10) 

	//in float64 after decimal point it gives only 5 digits
	var n11 float64 = 232.343445435443544343434
	fmt.Printf("%T %v\n",n11,n11)

	var a,b,c int8 = 22,34,45
	fmt.Println(a,b,c)


	

}
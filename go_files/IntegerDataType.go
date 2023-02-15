package main
/* there are three types in basic data type
1)numbers
2)booleans
3)strings
under numbers three sub categories
1)integers:signed integers(int) and unsigned integers(uint)*/

import "fmt"
func main(){
	//signed integers

	//int8 (8-bit signed integer whose range is -128 to 127)
	var a int8 = 127
	fmt.Println(a)

	//int16 (16-bit signed integer whose range is -32768 to 32767))
	var b int16 = 1000
	fmt.Println(b)

	//int32 (32-bit signed integer whose range is -2147483648 to 2147483647)
	//note:int is used for Both int and uint contain same size, either 32 or 64 bit.
	//rune:It is a synonym of int32 and also represent Unicode code points.
	var c int = 2146332323
	var d int32 = 232223424
	var s rune = 232333143
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(s)

	//int64 (64-bit signed integer whose range is -9223372036854775808 to 9223372036854775807)
	var e int64 = 2334143231341323234
	var f int = 3424241452413441123
	fmt.Println(e)
	fmt.Println(f)

	//unsigned integers

	//uint8	8-bit unsigned integer
	//uint8 (8-bit unsigned integer whose range is 0 to 255 )
	//byte	It is a synonym of uint8.
	var i uint8 = 33
	var j byte = 100
	fmt.Println(i)
	fmt.Println(j)

	//uint16 :16-bit unsigned integer
	//uint16 (16-bit unsigned integer whose range is 0 to 65535 )
	var k uint16 = 34636
	fmt.Println(k)

	//uint32:32-bit unsigned integer
	//uint32:(32-bit unsigned integer whose range is 0 to 4294967295)
	//uint:Both int and uint contain same size, either 32 or 64 bit.
	var h uint32 = 2324233232
	fmt.Println(h)
	var u uint = 3431353132
	fmt.Println(u)

	//uint64:64-bit unsigned integer
	//uint:Both int and uint contain same size, either 32 or 64 bit.
	//uint64 (64-bit unsigned integer whose range is 0 to 18446744073709551615 )
	var q uint64 = 3112314131312341231
	fmt.Println(q)
	var r uint = 2323342143141243232
	fmt.Println(r)




}
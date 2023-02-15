package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "100.23"
	var str2 string = "200"

	fmt.Println(strconv.ParseFloat(str1, 32))
	fmt.Println(strconv.ParseInt(str2, 2, 8))

	/*s : String value which is to be parsed in the integer number.
	base : The base value of the given value, it can be 0, 2 to 36.
	bitSize : It defines the integer type,
	0 for int
	8 for int8
	16 for int16
	32 for int32
	64 for int64 */

}

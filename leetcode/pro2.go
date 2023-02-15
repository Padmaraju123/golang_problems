package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	a, b, c := reader.ReadLine() //it will take the user string and return slice of ascii numbers and bool value and error

	fmt.Printf("%T %v\n", a, a)

	fmt.Printf("%T %v\n", b, b)

	fmt.Printf("%T %v\n", c, c)

	return_string := string(a)

	fmt.Printf("%T %v\n", return_string, return_string)

	sp := strings.Split(return_string, "")

	for _, v := range sp {
		fmt.Printf("%T %v\n", v, v)
	}

	tTemp,_:= strconv.ParseInt(readLine(reader), 10, 64)
	fmt.Printf("%T %v\n", tTemp, tTemp)
	t := int32(tTemp)
	fmt.Printf("%T %v\n", t, t)

}
	func readLine(reader *bufio.Reader) string {
		str, _,_ := reader.ReadLine()
		fmt.Printf("%T %v\n",str,str)
		return strings.TrimRight(string(str), "\r\n")
	}




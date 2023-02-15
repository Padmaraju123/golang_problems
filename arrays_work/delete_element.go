package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	out := scanner.Scan()       //it returns true allways

	fmt.Printf("%T %v\n", out, out)

	word := scanner.Text()

	bytes_slice := scanner.Bytes()         //ascii numbers of input nothing but bytes

	fmt.Printf("%T %v\n", word, word)

	fmt.Printf("%T %v\n", bytes_slice, bytes_slice)

	//convert byte_slice to string
	byteword := string(bytes_slice)

	fmt.Printf("%T %v\n", byteword,byteword)

	//convert word to int
	word_int,_ := strconv.Atoi(byteword)
	fmt.Printf("%T %v\n",word_int,word_int)
}

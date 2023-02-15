package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args) //to give the user inputs
	var jh = os.Args[1:]
	fmt.Println(jh)

	var nree = []int{}

	for _, k := range jh {
		fmt.Printf("%T\n", k)
		out, _ := strconv.Atoi(k)
		fmt.Printf("%T\n",out)
		nree = append(nree,out)
		fmt.Println(nree)
	}

	fmt.Println(nree)

	var wer = []string{}

	for _ ,nn := range nree{
		fmt.Printf("%T\n",nn)
		qwe := strconv.Itoa(nn)
		fmt.Printf("%T\n",qwe)
		wer = append(wer,qwe)
		fmt.Println(wer)
	}

	fmt.Println(wer)

}

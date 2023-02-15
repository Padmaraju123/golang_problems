package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	out, _, _ := reader.ReadLine()
	fmt.Println(out)
	for i,v := range out{
		if string(v)==" "{
			continue
		}
		fmt.Printf("%v %v\n",string(v),string(out[i]))
	}
}

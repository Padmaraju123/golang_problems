package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func summ(sli []string)int{
	sum := 0
	for _,v := range sli{
		vv,_:=strconv.Atoi(v)
		sum=sum+vv
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	out, _, _ := reader.ReadLine()
	sent := string(out)
	slice_str := strings.Split(sent, " ")
	fmt.Printf("%T %v\n", slice_str, slice_str)
	fmt.Println(summ(slice_str))

}

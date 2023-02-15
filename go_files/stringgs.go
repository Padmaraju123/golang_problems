package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	user_i := bufio.NewReader(os.Stdin)
	get_user,_ := user_i.ReadString('\n')


	//  res1 := strings.SplitAfter(str1, ",")

	called_one(get_user)

}

func called_one(word string){
	var sum_1 int
	outt := strings.SplitAfter(word,"!")
	fmt.Printf("%+v\n",outt)
	var le = len(outt)
	for i:=0;i<le;i++{
		tt,_:= strconv.Atoi(outt[i])
		fmt.Printf("%T %v\n",tt,tt)
		sum_1+=tt
	}

}

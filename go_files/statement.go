package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	i,err := strconv.Atoi("WEW")

	fmt.Println(i)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("NO ERROR")
	}

	//above if statement modification

	if j,erro := strconv.Atoi("FDFD") ; erro != nil{
		fmt.Println(erro)
	}else {
		fmt.Println(j)
	}

	if args := os.Args ; len(args) != 2{
		fmt.Println("ONE ARGUMENT IS REQUIRED")
	}else if km,er := strconv.Atoi(args[1]) ; er != nil {
		fmt.Println("The argument must be integer",er,km)
	}else{
		fmt.Println("NO NO")
	}

	g := os.Args
	fmt.Println(g)

}
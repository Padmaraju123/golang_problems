package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var mm = map[string]string{
		"ram":    "cricket",
		"naresh": "football",
		"vani":   "tennis",
		"rahim":  "cricket",
	}

	//addition key value pair
	df := bufio.NewReader(os.Stdin)
	sdd, _ := df.ReadString('\n')
	jj := strings.Split(sdd, " ")
	mm[jj[0]] = jj[1]
	fmt.Printf("%T %v\n", mm, mm)

	//update the key value
	jjhf,_ := df.ReadString('\n')
	ddd := strings.Split(jjhf," ")
    mm[ddd[0]]= ddd[1]
	fmt.Println(mm)

	//remove the key in maps
	delete(mm,"wee")

	fmt.Println(mm)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func checking(sli []byte)bool{
	mapp := make(map[string]bool)

	for _,v := range sli{
		if mapp[string(v)] != true{
			mapp[string(v)] = true
		}
	}
	
	fmt.Println(mapp)
	if len(mapp)==1{
		return true
	}else{
		return false
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	slice, _, _ := reader.ReadLine()

	fmt.Printf("%T %v\n", slice, slice) //gives slice of uint8 ascii numbers of alphabets

	//converting  uint8 slice of ascii numbers to string
	word := string(slice)
	fmt.Printf("%T %v\n",word,word)

	fmt.Println(checking(slice))
}

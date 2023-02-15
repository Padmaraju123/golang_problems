package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//opening a file
	my, err := os.OpenFile("third_file.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664) //to create a file with name
	if err != nil {
		//to terminate the program
		log.Fatal(err)
	}

	defer my.Close() //it will close the entire the file at end because we use defer statement

	//how to write the text inside the file
	byteslice := []byte("i am learning")

	bytewritten, err := my.Write(byteslice)
	if err != nil {
		//to terminate the program
		log.Fatal(err)
	}
	fmt.Println(bytewritten)
	fmt.Printf("no of bytes %d\n", bytewritten)

	//direct way to create a file and write 
	newbytes := []byte("go programming is cool")
	err = ioutil.WriteFile("fourth_file.txt",newbytes,0644)
	if err != nil {
		//to terminate the program
		log.Fatal(err)
	}
	fmt.Println(newbytes)
	fmt.Printf("no of bytes %d\n", newbytes)


}

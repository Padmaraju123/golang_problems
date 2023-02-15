package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//how to read the text in a file
	//open the file
	get, err := os.Open("iam.txt")
	//if error end the program
	if err != nil {
		log.Fatal(err)
	}
	//close the file after all operations
	defer get.Close()

	//create a slice with make function
	new_S := make([]byte, 3) //for this i can read exactly two bytes from a file

	//if the getting file byte size is smaller than our byte size it cause error
	numberbytes, err := io.ReadFull(get, new_S)
	if err != nil {
		log.Fatal("we have not enough bytes")
	}

	fmt.Printf("no of bytes can read %d\n", numberbytes)
	fmt.Printf("data is %s\n", new_S)

	//to read entire bytes of the file
	get, err = os.Open("file4.go")

	if err != nil {
		log.Fatal(err)
	}

	defer get.Close()

	fulldata, err := ioutil.ReadAll(get)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data as string %s\n",fulldata)

	//how to read entire text in file 
	fulldata,err = ioutil.ReadFile("fifth.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n",fulldata)


}



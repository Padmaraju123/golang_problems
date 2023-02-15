package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//how to create a file in go

	var newfile *os.File
	//here *os.file is used to store the address of the newfile in memory
	fmt.Printf("%T\n", newfile)

	//create a file
	var err error                              //this is error type which gives error
	newfile, err = os.Create("first_file.txt") //if you want error put filename as (//filename.txt)

	if err != nil {
		//how to terminate the code if error occur
		//os.Exit(1)
		//another way to terminate the code
		log.Fatal(err) //it specifies the time when compare to os.exit
	}

	//how to clear/truncate the text in the file
	os.Truncate("first_file.txt", 0) //0 means to clear all text

	//how to close the file
	newfile.Close()

	//how to open a file
	file, err := os.Open("first_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	//how to get name of the file on console
	var file_name os.FileInfo
	file_name, err = os.Stat("first_file.txt")
	fmt.Println(file_name.Name())
	fmt.Println("size in bytes: ", file_name.Size())
	fmt.Println("last modified: ", file_name.ModTime())
	fmt.Println("Is Directory: ", file_name.IsDir())
	fmt.Println("Permissions: ", file_name.Mode())

	//how to check the file exist or not
	check1, err := os.Stat("second_file.go")
	_ = check1
	fmt.Println(os.IsNotExist(err))

	if err != nil {
		if os.IsNotExist(err) { //it returns error
			fmt.Println("This file doesnot exist!!!")
		}
	}

	//how to rename the file
	os.Rename("first_file.txt", "second_file.txt")

	//how to remove a file
	os.Remove("second_file.txt")

}

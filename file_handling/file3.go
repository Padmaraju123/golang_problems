package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//another way to create afile and write on it
	file, err := os.OpenFile("fifth.txt", os.O_WRONLY|os.O_CREATE, 0664)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bufferWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("bytes written %d\n", bytesWritten)

	//how to find no of bytes left in bufferwriter
	left_bytes := bufferWriter.Available()
	fmt.Println(left_bytes)

	//how to check how many are used
	used_bytes := bufferWriter.Buffered()
	fmt.Println(used_bytes)

	//how to write
	added, err := bufferWriter.WriteString("\ndfhd adfhksfnks")
	_ = added

	//here what ever write it does not stored in our file

	used_bytes = bufferWriter.Buffered()
	fmt.Println(used_bytes)

	//how to add in my file
	bufferWriter.Flush()

}

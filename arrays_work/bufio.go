package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//open a file located in directory
	file, err := os.Open("description.txt")
	fmt.Printf("%T %v\n", file, file)
	fmt.Printf("%T %v\n", err, err)

	//read lines in description.txt file from bufio package with NewReader
	reader := bufio.NewReaderSize(file,1024*1024)

	//peek method reads integer number how many characters to read in description file and gives slices of bytes
	//it returns uint8 data type because //byte is alias of uint8
	slice_bytes_numbers, _ := reader.Peek(6)
	fmt.Printf("%T %v\n", slice_bytes_numbers, slice_bytes_numbers)

	//how to convert slice_of_bytes to string
	word := string(slice_bytes_numbers)
	fmt.Printf("%T %v\n", word, word)

	//Read method from reader takes slice 0f bytes and returns number of bytes read by peek method
	num_of_bytes, _ := reader.Read(slice_bytes_numbers)
	fmt.Printf("%T %v\n", num_of_bytes, num_of_bytes)

	//read a single byte in given description file
	byte, _ := reader.ReadByte()
	fmt.Printf("%T %v\n", byte, byte)

	//convert byte to string
	//here we get y because the reader reads next character in description file
	char := string(byte)
	fmt.Printf("%T %v\n", char, char)

	//ReadBytes method is used to read the bytes unto given argument in method
	slic_byte, _ := reader.ReadBytes('\n')
	fmt.Printf("%T %v\n", slic_byte, slic_byte)
	fmt.Printf("%T %v\n",string(slic_byte),string(slic_byte))

	//ReadLine method is used to take the long sentence with spaces upto next line and convert into slice of bytes 
	sli_bytes,_,_:=reader.ReadLine()
	fmt.Printf("%T %v\n",sli_bytes,sli_bytes)

	//comvert it into sentence
	fmt.Println(string(sli_bytes))
	
	//ReadRune reads one character and return and int32/rune number 
	sli_rune,size,_:=reader.ReadRune()
	fmt.Printf("%T %v %v\n",sli_rune,sli_rune,size)
	fmt.Printf("%T %v\n",string(sli_rune),string(sli_rune))


	//read the sentence upto delimit and return slice of bytes
	s_bytes,_:=reader.ReadSlice(' ')
	fmt.Printf("%T %v\n",s_bytes,s_bytes)
	fmt.Println(string(s_bytes))

	//read the sentence and returns as it is
	word1,_ := reader.ReadString(' ')
	fmt.Printf("%T %v\n",word1,word1)

	//it returns no of bytes till now
	count:=reader.Buffered()
	fmt.Printf("%T %v\n",count,count)

	//it will discard the desired bytes
	value,_:=reader.Discard(4)
	fmt.Printf("%T %v\n",value,value)

	//now i am reading after the discarded
	word2,_ := reader.ReadString(' ')
	fmt.Printf("%T %v\n",word2,word2)

	//it will returns size of underlying 
	fmt.Printf("%T %v\n",reader.Size(),reader.Size())



}

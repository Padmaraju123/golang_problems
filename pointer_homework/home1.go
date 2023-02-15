package main
import "fmt"

func main(){
	var s = 100
	add_s := &s
	fmt.Println(add_s)
	fmt.Println(&s)
	fmt.Printf("the data type of s %T, and the value of s %v\n",s,s)
	fmt.Printf("the data type of add_s %T , and the value of add_s %p\n",add_s,add_s)
	fmt.Printf("the data type of &s %T , and the value of &s %p\n",&s,&s)

	add_of_add_S := &add_s
	fmt.Println(add_of_add_S)

	//how to intialise pointer to particular data type
	var ptr1 *int
	fmt.Printf("the data type of ptr1 is %T , and the value/address is %v\n",ptr1,ptr1)

	/*//to use the ptr for address to the int data types values only
	num := 100
	ptr1 = &num
	fmt.Printf("the data type of num %T, and the value is %v\n",num,num)
	fmt.Printf("the data type of ptr1 %T, and the address/value is %p\n",ptr1,ptr1)

	//another way to fix the pointer to data type
	ptr2 := new(string)
	word := "golang"
	ptr2 = &word
	fmt.Println(ptr2)

	//to change the value of variable if it has pointer of that variable
	*ptr2 = "golang is very bad"
	fmt.Println(*ptr2)
	fmt.Println(word)*/

}
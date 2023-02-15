package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//backing or underlying array
	//when we create a slice behind the scenes the GO creates hidden array it is backing array
	//backing array stores the elements not the slice
	//then what elements are stored by slice how to find backing array by slice
	//Go implements a slice as a data structure called slice header
	//the slice header contains three fields
	//1)address of the backing array(pointer)
	//2)the length of the slice not all the elements of backing array should belong to the slice
	//3)the capacity of the slice  this is size of the backing array after the slice first element
	//the nil slice does not have backing array

	//slice expression doesn't create a new backing array
	//the slice returned by a slice expression refers to the backing array of the original slice
	//when changing a slice obtained from a slicing expression the backing array is changed and all othe sliced creating from the same backing
	//array

	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(len(s1))

	//by using slice expression creating new slices
	s2, s3 := s1[0:2], s1[1:3]
	//now update the value of s3 it will effect to all slices because the backing array is same to all
	s3[1] = 600
	fmt.Println(s1, s2, s3)

	//when a slice is created by slicing an array , the array becomes the backing array of the new slice
	var c_Arr = [4]int{100, 200, 300, 400}
	slice1, slice2 := c_Arr[0:2], c_Arr[1:3]

	c_Arr[1] = 1000
	fmt.Println(c_Arr, slice1, slice2)

	//to avoid this situation we use append method

	var cars = []string{"Ford", "honda", "Audi", "Range Rover"}

	var k = []string{}

	k = append(k, cars[1:3]...)

	fmt.Println(k)

	//if you change it wont change because different backing array
	cars[1] = "hero"

	fmt.Println(cars, k)

	//the length and capacity of the backing array

	var chh = []int{323, 4343, 6454, 5656, 2323232}
	var fg = chh[2:3]

	//length is the no of elements in substring
	//capacity is no of elements from substring first index of the backing array

	fmt.Println(len(fg), cap(fg))

	//address of the backing array

	fmt.Printf("%p\n", &chh)

	//address are different but the backing array is same
	fmt.Printf("%p %p\n", &chh, &fg)

	//how to find size of the memory of array and slice
	m_A := [5]int{1, 2, 3, 4, 5}
	m_s := []int{1, 2, 3, 4, 5}
	//note: the memory take by slice is less than array
	fmt.Printf("%d\n", unsafe.Sizeof(m_A))
	fmt.Printf("%d\n", unsafe.Sizeof(m_s))

	//creating a slice
	type name []string

	ok := name{"raju", "indra"}
	fmt.Println(ok[1])

}

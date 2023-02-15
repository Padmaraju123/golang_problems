package main

import "fmt"

func main(){
	//creating nested array
	nested_array := [2][3]int{
		{10,20},

	}
	fmt.Printf("%T %v\n",nested_array,nested_array)

	//assigning values to nested_Array
	nested_array[0][2] = 30

	fmt.Println(nested_array)

	//assigning values to nested_Arry using for loop
	for i:=0 ; i<3;i++{
		var w int
		fmt.Scanln(&w)
		nested_array[1][i]= w
	}

	fmt.Println(nested_array)

	//assining values to nested_Arr using for loop with range
	for i,v := range nested_array[1]{
		nested_array[0][i] = v
	}

	fmt.Println(nested_array)

	for a := 0 ;a<len(nested_array);a++{
		for b:=0;b<len(nested_array[0]);b++{
			fmt.Println(nested_array[a][b])
		}
	}


	var arr1 []string
	fmt.Printf("%T %v\n",arr1,arr1)

	var arr2 [2]string
	fmt.Printf("%T %v\n",arr2,arr2)

	var arr3 []int
	fmt.Printf("%T %v\n",arr3,arr3)

	var arr4 [3]int
	fmt.Printf("%T %v\n",arr4,arr4)

	var arr5 = [4]int{}
	fmt.Printf("%T %v\n",arr5,arr5)

}
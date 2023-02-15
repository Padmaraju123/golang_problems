package main

import (
	"fmt"
)

func main(){
	var arryys [2]int 
	fmt.Printf("%v %T\n",arryys,arryys)
	fmt.Printf("%#v\n",arryys)
	arryys = [2]int{23,43}
	fmt.Println(arryys)

	vere := [3]int{100,0202,2323}  //only integers are used in this case
	fmt.Println(vere)

	var hdkfd = [4]string{"23","5","4","4"}
	fmt.Println(hdkfd)
	fmt.Println(hdkfd[0]) //Access the element

	var hdkf = [4]string{"23","5"}  //here we didn't assign all values
	fmt.Println(hdkf)


	var a1 = [2]float32{}
	fmt.Printf("%#v\n",a1)
	fmt.Println(a1)

	//ellipsis operator
	var a2 = [...]int{23,4,23,4,23,23,2}   //here we can assign desired elements
	fmt.Println(a2)
	fmt.Printf("length of a2 %d\n",len(a2))
	fmt.Println("LENGTH OF A2",len(a2))

	a2[0] = 100  //updating the value in arrays we can't delete the elements because it has fixed length
	fmt.Println(a2)

	//first way to iterate the array
	for i,hh := range a2 {

		fmt.Printf("THE INDEX %d the value %d\n",i,hh)
		fmt.Printf("THE VALUE IS %d\n",a2[i])

	}

	//second way to iterate the array

	for i:=0 ; i<len(a2); i++{
		fmt.Printf("THE VV IS %d\n",a2[i])
	}
	

	//two-dimensional array

	var dim = [3][3]int{
		{1,2,3},
		{4,5,6},
		{1000,2000,3000},
	}

	fmt.Println(dim)

	//declaring arrays
	//first way
	var arr1 = [2]int{102,343}
	fmt.Println(arr1)
	var arr2 = [...]string{"raju","r"}
	fmt.Println(arr2)

	arr3 := [3]float32{2.34,34.34}
	fmt.Println(arr3)

	var arr4 = [4]int{0:101,3:104}   //it will assing values on index position
	fmt.Println(arr4)

	var le = len(arr4)

	for i := 0 ;i<le ;i++{
		fmt.Println(arr4[i])
	}



}
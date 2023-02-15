package main

import "fmt"

func add(n int)[]int{

	arr1 := make([]int ,n)

	arr2 := make([]int ,n)
	

	for i:=0;i<n;i++{
		fmt.Scan(&arr1[i])
		
	}

	for j:=0;j<n;j++{
		fmt.Scan(&arr2[j])
		
	}

	var final = []int{}

	for k:=0;k<n;k++{
		final = append(final,arr1[k]+arr2[k])
	}

	return final
}

func main(){
	var n int
	fmt.Scanln(&n)
	fmt.Println(add(n))

}
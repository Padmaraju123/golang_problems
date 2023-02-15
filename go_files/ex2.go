
package main

import "fmt"

func unique(arr []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for e := range arr {

		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}

	return result 
}

func main() {
	array1 := []int{1, 5, 3, 4, 1, 6, 6, 6, 8, 7, 13, 5}
	fmt.Println(array1)
	unique_items := unique(array1)
	fmt.Println(unique_items)
}

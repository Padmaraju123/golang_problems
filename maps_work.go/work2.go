// Golang program to remove duplicate
// values from Slice
package main

import (
	"fmt"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	// Assigning values to the slice
	intSliceValues := []int{1,2,3,4,5,2,3,5,7,9,6,7}

	// Printing original value of slice
	fmt.Println(intSliceValues)

	// Calling function where we
	// are removing the duplicates
	removeDuplicateValuesSlice := removeDuplicateValues(intSliceValues)

	// Printing the filtered slice
	// without duplicates
	fmt.Println(removeDuplicateValuesSlice)
}

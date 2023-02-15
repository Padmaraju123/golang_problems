package main

import (
	"fmt"
)

func duplicate(sli []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range sli {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func main() {

	intSliceValues := []int{1, 2, 3, 4, 5, 2, 3, 5, 7, 9, 6, 7}

	fmt.Println(intSliceValues)

	out := duplicate(intSliceValues)

	fmt.Println(out)

}

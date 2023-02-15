package main

import "fmt"

//here we are doing for maps and slices

func changeSlice(s []int){
	for i := range s{
		s[i]++
	}

}

func changeMaps(m map[string]int){
	m["A"] = 100
	m["B"] = 200
	m["C"] = 300
 
}

func main(){
	var slice = []int{10,20,30}
	changeSlice(slice)
	fmt.Println(slice)

	mymap := map[string]int{
		"A":10,
		"B":20,
		"C":30,
	}
	changeMaps(mymap)
	fmt.Println(mymap)
}
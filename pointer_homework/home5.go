package main

import "fmt"

//passing maps to the functions

func best(m *family){
	(*m)["raju"] = 1000
	(*m)["sandhya"] = 2000
	(*m)["purushotham"]= 50000


}

func normal(m map[string]int){
	m["raju"] = 200
	m["sandhya"] = 23232
	m["purushotham"] = 223523

}

//create a map
type family map[string]int

func main(){
	mm:=family{
		"raju":23,
		"sandhya":40,
		"purushotham":2323,
	}

	normal(mm)
	fmt.Println(mm)
	//here it is changing with normal also

	//ofcourse we do pointer way also
	best(&mm)
	fmt.Println(mm)
}
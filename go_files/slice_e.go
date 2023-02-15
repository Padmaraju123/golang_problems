package main
import "fmt"

func main(){
	var sli []int   //unintialsed
	fmt.Println(sli == nil)
	fmt.Printf("%#v\n",sli)

	var slii = []int{} 
	fmt.Println(slii == nil)
	fmt.Printf("%#v\n",slii)

	//we can't compare the slices with == operator

	s1,s2 := []int{1,2,3} , []int{1,2,3}

	//fmt.Println(s1==s2)      //error

	//to compare the slices we go with for loop compare element by element

	var eq = true

	s1 = nil   //for this statement the for skipped

	for i , val := range s1{
		if val != s2[i]{
			eq = false
			break
		}
	}

	//to avoid this problem compare the lenths
	if len(s1) != len(s2){
		eq = false
	}

	if eq{
		fmt.Println("these slices are equal")
	}else{
		fmt.Println("slices are not equal")
		}

}
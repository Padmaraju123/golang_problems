package main

import "fmt"

func get_small_num(sli []int)int{
	check := sli[0]
	for _,v := range sli{
		if v < check{
			check = v
		}	
	}
	return check
}


func main(){
	var num int
	fmt.Scanln(&num)

	sli := make([]int , num)

	for i:=0;i<num;i++{
		fmt.Scanln(&sli[i])
	}

	fmt.Printf("%T %v\n",sli,sli)

	fmt.Println(get_small_num(sli))
}
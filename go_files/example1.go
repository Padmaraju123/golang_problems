package main
import (
	"fmt"
	"os"
	"strconv"
)
func removeDup(arry ...string){

	ac_arr := []int{}

	for j:= range arry{
		gg ,_ := strconv.Atoi(arry[j])
		ac_arr=append(ac_arr, gg)
	}

	map_v := map[int]bool{}
	result := []int{}

	for i:=range ac_arr{
		if map_v[ac_arr[i]]==false{
			map_v[ac_arr[i]] = true
			result= append(result, ac_arr[i])
		}
	}

	fmt.Println(result)

}

func main(){
	var user_arr = os.Args
	var out_arr = user_arr[1:]
	fmt.Println(out_arr)
	removeDup(out_arr...)

}
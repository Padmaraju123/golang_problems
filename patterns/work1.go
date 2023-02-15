package main 

import "fmt"


func get(n int,dd [3][3]int)[3][3]int{

	for a:=0;a<3;a++{
		le := len(dd)
		dd[a][le-1]=n
	}

	return dd


}

func main(){
	var d_arr = [3][3]int{}
	
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			var out int
			fmt.Scanln(&out)
			d_arr[i][j] = out
		}
	}

	var n int
	fmt.Scanln(&n)
	fmt.Println(get(n,d_arr))
}
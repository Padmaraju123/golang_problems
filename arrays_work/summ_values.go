package main


import "fmt"

func pairs(sli []int,k int)[][]int{
	le := len(sli)
	var out [][]int
	for i:=0;i<le-1;i++{
		for j:=i+1;j<le;j++{
			if sli[i]+sli[j]==k{
				j := [][]int{{sli[i],sli[j]}}
				out = append(out,j...)
			}
		}
	}
	return out
}

func main(){
	sli := []int{1,2,3,4,5}
	k := 5
	fmt.Println(pairs(sli,k))
	
}
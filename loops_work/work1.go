package main
import "fmt"

func ll(m,n int)[]int{
	var slr = []int{}
	for i:=m;i<=n;i++{
		slr = append(slr,i)
	}
	return slr
}

func main(){
	var s1,s2 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Println(ll(s1,s2))

}
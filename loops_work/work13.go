package main
import "fmt"

func pow(m,n int)int{
	var ss = 1
	for i:=1;i<n+1;i++{
		ss*=m
	}
	return ss
}

func main(){
	var m,n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Println(pow(m,n))
}
package main
import "fmt"

func vao(m,n int)int{
	var ot = 1

	var i = m
	for i<n+1{
		ot*=i
		i++
	}
	return ot
}

func main(){
	var m,n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Println(vao(m,n))
}
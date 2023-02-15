package main
import (
	"fmt"
)	

func ghh(m,n int)int{
	var sss = 0
	for i:=1;i<n+1;i++{
		sss+=i**i

	}
	return sss
}

func main(){
	var m,n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)

}
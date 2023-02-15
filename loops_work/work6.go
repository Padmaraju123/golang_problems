package main
import "fmt"

func rev(n int){
	 k := n
	for i:=1;i<k+1;i++{
		fmt.Println(n)
		n--

	}
}

func main(){
	var n int
	fmt.Scanln(&n)
	rev(n)
}
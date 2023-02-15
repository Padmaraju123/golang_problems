package main
import "fmt"

func change(a *int){
	*a = 200

}

func unchange(b int){
	b = 300
}

func main(){
	x := 100
	p := &x
	fmt.Println(x) 
	change(p)
	fmt.Println(x)    //here we don't return a value but change due the address or pointer 

	//if you pass normal value it wont work
	fmt.Println(x)
	unchange(x)
	fmt.Println(x)
}
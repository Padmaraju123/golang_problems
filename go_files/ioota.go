package main
import "fmt"

func main(){
	const (
		x=iota 
		y=iota 
		z=iota
	)

	fmt.Println(x,y,z)

	const (
		a=iota
		b 
		c
	)

	fmt.Println(a,b,c)

	const (
		i = iota+1
		j
		k
	)

	fmt.Println(i,j,k)

	const (
		jj= iota 
		_						//skipping the identifier

		kk


	)
	fmt.Println(jj,kk)

    const p = 7           
    const q float64 = 3.1
    fmt.Println(p * q)

	const (
        min1 = "raju"
        max1
        max2 
    )
 
    fmt.Printf("%T %v\n", max1, max2)
}
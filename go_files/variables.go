package main
import "fmt"
var ss = 2323
func main(){

	fmt.Println(ss)

	var df int
	fmt.Println(df)   //it prints 0 value ; if float also 0; if string ""(empty space); if boolean false 

	a := "ctfyhc"
	fmt.Printf("%T\n",a)

	i,j := "sdsd",1212
	fmt.Printf("%q %v\n",i,j)

	i , j = "Sdsd",54543
	fmt.Printf("%q %v\n",i,j)


	var r , s  int
	r = 2323
	s = 2323
	fmt.Println(r,s)

	var (
		name string = "raju"
		age int = 23
		points float32 = 20.5
		gender string
	)

	gender = "male"

	fmt.Println(name,age,points,gender)

	var er ,erer = 23 ,423
	fmt.Printf("%T\n",er)
	fmt.Printf("%T\n",erer)

	er = erer    //it is not occur error because same data type

	var ek int
	var ddd string
	var dffgf bool
	var dndfd float32
	var dffg float64
	fmt.Println(ek,ddd,dffg,dndfd,dffgf)


	

	/*_ = a              //if don't use a varible then assing like this
	var b = "raju"
	_ = b
	a = 1000
	name,age := "raju",23
	name ,age = "rj",25 
	_,_=name,age 
	var check = true
	fmt.Println(check)

	var (
		salary float64
		firstname string
		gender bool
	)

	fmt.Println(salary,firstname,gender)

	var x,y,z int 
	fmt.Println(x,y,z)

	var i,j int
	i,j=10,20  
	j,i=i,j        //swapping varibles

	fmt.Println(i,j)*/

}
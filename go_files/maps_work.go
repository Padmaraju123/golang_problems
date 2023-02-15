package main

import "fmt"

func main(){
	//creating a empty unintialised map
	var m_1 map[int]int
	fmt.Printf("%#v\n",m_1)
	fmt.Printf("%d\n",m_1)

	//creating a empty intialised map
	var m_2 = map[int]int{}
	fmt.Printf("%#v\n",m_2)
	fmt.Printf("%v\n",m_2)

	//we can't use slice as key
	//we can't append key value pairs for unintialised maps but initialised maps we can.
	m_2[23] = 232
	fmt.Println(m_2)
	fmt.Printf("the key value is: %d\n",m_2[23])
	
	//how to create a nil map/arry/slice with make function but initialised 
	ar1 := make([]int,3)
	fmt.Println(ar1)
	ar1[0] = 123
	fmt.Printf("%#v\n",ar1)

	slii := make([]string ,2,3)
	fmt.Println(slii)
	slii[0]="READ"
	slii[1]="DFDS"
	fmt.Println(slii)

	maap := make(map[int]string)

	fmt.Printf("%v\n",maap)
	maap[545] = "rerere"
	fmt.Println(maap)

	//create a direct map
	var msd = map[int]string{
	23:"wdsd",
	23534:"33434",
	}
	fmt.Println(msd)
	//we can update values of the keys
	msd[23] = "adfaa"
	fmt.Println(msd)
	//print the values of the map
	//if the key occur then it provides the value of that key
	//if not it gives empty string of string value 
	//if not it gives 0 value of int value
	fmt.Printf("the value of the key is %s\n",msd[3434])
	fmt.Printf("the value of the key is %s\n",msd[23])

	//how to check whether particular key-value of the map

	cc ,boo  := msd[2233]
	if boo{
		fmt.Println(cc)
	}else{
		fmt.Println("not exist")
	}

	//iterate the maps
	for wee ,wea := range msd{
		fmt.Println(wee,wea)
	}

	

}




	


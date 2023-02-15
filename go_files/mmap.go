package main

import "fmt"

func main(){
	var mm map[string]int
	fmt.Printf("%#v\n",mm)
	fmt.Printf("%v\n",mm)
	fmt.Println(len(mm))
	fmt.Printf("the first key %d\n",mm["edsd"])  //%d is used because int value 
	var kk map[int]string
	fmt.Printf("the fisrt ket %q\n",kk[23])   //%q is used for quoted strings values

	//we don't use slice as key in map because slice is not comparable but we use arrays as keys 

	//var kk map[[]int]string

	//assign the key value pairs in map
	//mm["raju"] = 1000
	//mm["ravi"] = -122

	//fmt.Printf("the mm map is %#v\n",mm)  it ours error because it is nil map before
	//it is not good assign values with initialise

	wewe := map[string]float32{}   //now it initialised so we can add key value pairs

	wewe["raju"] = 23.34
	wewe["erer"] = 23.323

	fmt.Printf("the map is %#v\n",wewe)

	//we can create nil map by make function
	mapp := make(map[int]string)
	mapp[1] = "raaa"
	fmt.Printf("%#v\n",mapp)

	//we can map directly 
	mapp2 := map[string]float32{"rere":23.4,"ddre":23.423,"dsds":23232.3} 

	fmt.Printf("%#v\n",mapp2)

	//To replace the value of the key
	mapp2["rere"] = 23.445

	fmt.Println(mapp2)

	//if you print the value of the map key
	//if exist it returns the value of the particular key
	fmt.Println(mapp2["rere"])
	//if not exist it returns the zero
	fmt.Println(mapp2["SDE"])

	//how to check whether particular key-value pair exist
	va , bool_val  := mapp2["ree"]

	if bool_val{
		fmt.Printf("the value of the key \"rere\" is : %v\n",va)
	}else{
		fmt.Printf("the value doesn't exist to this key \"rere\"\n")
	}

	//iterating of maps
	for k,v := range mapp2{
		fmt.Printf("the key: %#v,the value: %#v\n",k,v)
	}

	//to delete the map element
	delete(mapp2,"rere")
	fmt.Printf("%#v\n",mapp2)

	//we can't compare map with map but compare with nil or not.

	m11 := map[string]string{"23":"@323","34":"2323"}
	m22 := map[string]string{"23":"@323","34":"2323"}

	//firstly convert into maps into strings
	s11 := fmt.Sprintf("%s",m11)
	s22 := fmt.Sprintf("%s",m22)

	fmt.Printf("%#v %T\n",s11,s11)
	fmt.Printf("%#v %T\n",s22,s22)

	if s11 == s22{
		fmt.Printf("%v and %v are equal maps\n",s11,s22)
	}else{
		fmt.Printf("%v and %v are not equal maps\n",s11,s22)
	}


	//how to copy map to another map
	var rm = map[int]string{20:"dere",34:"sdee"}
	var mr = rm

	fmt.Println(mr)
	//if you change the element in rm it will reflect to mr because they are located in same memory location
	rm[20]= "eeere"
	fmt.Println(mr)

	//if you change the element in mr it will also same 
	mr[34] =" padasa"
	fmt.Println(rm)

	//creating a nil map by make function
	var new_mm = make(map[int]string)

	//by iterating to copy the map elements into another map 
	for b , c := range rm{
		fmt.Println(b,c)
		new_mm[b] = c
	}

	fmt.Printf("%#v\n",new_mm)

	//if you change the element for copied one it doesn't effect the original one
	new_mm[20] = "fdsfwer"
	fmt.Println(rm)
	fmt.Println(new_mm)

}
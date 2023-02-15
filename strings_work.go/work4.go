package main

import (
	"fmt"
)

func main() {
	//byte is alias of uint8
	stringBytes := []byte{72, 101, 108, 108, 111}

	fmt.Printf("%T %v\n",stringBytes,stringBytes)
	
	stringCreated := string(stringBytes)
	fmt.Printf("%T %v\n",stringCreated,stringCreated)

	for i,v := range stringCreated{
		fmt.Printf("%T %v\n\n\n",stringCreated[i],stringCreated[i])
		fmt.Printf("%T %v\n\n\n",v,v)


		fmt.Printf("%T %c\n\n\n",stringCreated[i],stringCreated[i])
		fmt.Printf("%T %c\n\n\n",v,v)

		fmt.Printf("%T %v\n\n\n",stringCreated[i],string(stringCreated[i]))
		fmt.Printf("%T %v\n\n\n",v,string(v))
	}
}

package main

import (
	"fmt"
)

func main() {

	/*user := bufio.NewReader(os.Stdin)
	get, _ := user.ReadString('\n')

	split_one := strings.Split(get, " ")

	fmt.Println(split_one)*/

	var slic = []int{}

	for i := 0; i < 3; i++ {
		var kk int
		fmt.Scanln(&kk)
		slic = append(slic, kk)
	}
	fmt.Println(slic)

}

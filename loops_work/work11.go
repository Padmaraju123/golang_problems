package main

import "fmt"

func av(dff int) float32 {
	df := float32(dff)
	return df / 10
}

func seu(n int) (int, float32) {
	var ss = 0

	for i := 0; i < 10; i++ {
		var dd int
		fmt.Scanln(&dd)
		ss += dd
	}
	hg := av(ss)

	return ss, hg

}

func main() {
	var n = 10
	out1, out2 := seu(n)
	fmt.Printf("%T %v and %T %v\n", out1, out1, out2, out2)

}

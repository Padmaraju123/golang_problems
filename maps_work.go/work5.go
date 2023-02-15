package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gett(n int,sr map[string]string) map[string]string {
	gh := bufio.NewReader(os.Stdin)

	var i = 0
	for i < n {
		ff,_:= gh.ReadString('\n')
		dfjj := strings.Split(ff, " ")
		sr[dfjj[0]] = dfjj[1]
		i++
	}

	return sr
}

func main() {
	var werr = map[string]string{
			"Ram": "Cricket",
			"Naresh": "Football",
			"Vani": "Tennis",
			"Rahim": "Cricket",
	}

	var num int
	fmt.Scanln(&num)
	fmt.Println(gett(num,werr))

}

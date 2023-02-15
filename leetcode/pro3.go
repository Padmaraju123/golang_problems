package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("  Hello, world  ", " "))
	fmt.Println(strings.TrimRight("Hello,   world  ", " "))
	fmt.Println(strings.TrimRight("...Hello, world...", "."))
	fmt.Println(strings.TrimRight("###Hello, world!!!", "!"))
	fmt.Println(strings.TrimRight("###Hello, world!!!", "!!!"))
	fmt.Println(strings.TrimRight("###Hello, world!!!", "#!"))
	fmt.Println(strings.TrimRight("sdsddsd\n\rffre","\n"))
}

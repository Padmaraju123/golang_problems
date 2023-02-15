package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var amount int

func ccount(int_S []int) {
	fmt.Scan(&amount)
	for _, v := range int_S {
		fmt.Println(v, "notes are", amount/v)
		amount = amount%v
	}

}

func count_total_notes(split_s []string) {
	int_S := []int{}
	for _, v := range split_s {
		k, _ := strconv.Atoi(v)
		int_S = append(int_S, k)
	}
	ccount(int_S)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	sli_b, _, _ := reader.ReadLine()
	sent := string(sli_b)
	split_s := strings.Split(sent, " ")
	count_total_notes(split_s)
}

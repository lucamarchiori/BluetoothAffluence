package main

import (
	"fmt"
	"math/rand"
)

var list []int

func addInt(list *[]int, n int) {
	if len(*list) == 0 {
		*list = append(*list, n)
	}

	for j := range *list {
		if (*list)[j] == n {
			break
		}
		if j == len(*list)-1 {
			*list = append(*list, n)
		}
	}
}

func main() {
	fmt.Println("Hello")

	min := 0
	max := 10
	for i := 0; i < 200; i++ {
		n := rand.Intn(max-min) + min
		fmt.Println(n)
		addInt(&list, n)

	}

	fmt.Println("Final: ")
	for i := range list {
		fmt.Println(list[i])
	}

}

package main

import "fmt"

func main() {
	var str string
	str = "Yogurt is a variant of milk"

	// OP: 89 111 103 117 ...
	for _, i := range str {
		fmt.Println(i)
	}
}

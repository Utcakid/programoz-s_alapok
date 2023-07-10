package main

import (
	"fmt"
)

func add(t, y int) int {
	return t + y
}

func main() {
	fmt.Println(add(42, 13))
}

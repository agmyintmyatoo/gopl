package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fib(n))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

package main

import "fmt"

func fibonachi(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonachi(n-1) + fibonachi(n-2)
	}
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(fibonachi(i), " ")
	}
}

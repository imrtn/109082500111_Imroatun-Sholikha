package main

import "fmt"

func pola(n int) {
	if n == 1 {
		fmt.Print(1, " ")
		return
	}

	fmt.Print(n, " ")
	pola(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scan(&n)

	pola(n)
}
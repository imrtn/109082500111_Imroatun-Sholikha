package main

import "fmt"

func main() {
	var suara [21]int
	var x int

	totalMasuk := 0
	suaraSah := 0

	fmt.Scan(&x)

	for x != 0 {

		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}

		fmt.Scan(&x)
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

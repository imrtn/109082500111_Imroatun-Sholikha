package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var a, b int
	i := 1

	fmt.Print("Pertandingan ", i, ": ")
	fmt.Scan(&a, &b)

	for a >= 0 && b >= 0 {
		fmt.Print("Hasil ", i, ": ")

		if a > b {
			fmt.Println(klubA)
		} else if b > a {
			fmt.Println(klubB)
		} else {
			fmt.Println("Draw")
		}

		i++
		fmt.Print("Pertandingan ", i, ": ")
		fmt.Scan(&a, &b)
	}

	fmt.Println("Pertandingan selesai")
}

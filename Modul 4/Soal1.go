package main

import "fmt"

func faktorial(n int, h *int) {
	*h = 1
	for i := 1; i <= n; i++ {
		*h *= i
	}
}

func permutasi(n, r int, h *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*h = fn / fnr
}

func kombinasi(n, r int, h *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*h = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var p1, k1, p2, k2 int

	fmt.Scan(&a, &b, &c, &d)

	permutasi(a, c, &p1)
	kombinasi(a, c, &k1)

	permutasi(b, d, &p2)
	kombinasi(b, d, &k2)

	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}

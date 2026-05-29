package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64
	var totalWadah []float64
	var rata float64

	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan kapasitas wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var jumlah float64
	var count int

	for i := 0; i < x; i++ {
		jumlah += ikan[i]
		count++

		if count == y || i == x-1 {
			totalWadah = append(totalWadah, jumlah)
			jumlah = 0
			count = 0
		}
	}

	fmt.Println("Total berat tiap wadah:")
	var totalSemua float64

	for i := 0; i < len(totalWadah); i++ {
		fmt.Printf("Wadah %d = %.2f\n", i+1, totalWadah[i])
		totalSemua += totalWadah[i]
	}

	rata = totalSemua / float64(len(totalWadah))

	fmt.Printf("Rata-rata berat tiap wadah: %.2f\n", rata)
}
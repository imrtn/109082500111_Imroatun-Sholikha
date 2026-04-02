package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktuSoal int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktuSoal)

		if waktuSoal != 301 {
			*soal++
			*skor += waktuSoal
		}
	}
}

func main() {
	var namaPeserta, namaPemenang string
	var soal, skor int
	var soalTerbanyak = -1
	var skorTerbaik = 999999

	fmt.Scan(&namaPeserta)

	for namaPeserta != "Selesai" {
		hitungSkor(&soal, &skor)

		if soal > soalTerbanyak || (soal == soalTerbanyak && skor < skorTerbaik) {
			soalTerbanyak = soal
			skorTerbaik = skor
			namaPemenang = namaPeserta
		}

		fmt.Scan(&namaPeserta)
	}

	fmt.Println(namaPemenang, soalTerbanyak, skorTerbaik)
}

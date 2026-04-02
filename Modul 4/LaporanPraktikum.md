# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
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

```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal1.png)

[Program ini meminta pengguna untuk memasukkan empat bilangan bulat yaitu a, b, c, dan d. Setelah itu, program akan menghitung nilai permutasi dan kombinasi dari pasangan (a, c) serta (b, d) dengan memanfaatkan prosedur faktorial.

Perhitungan dilakukan menggunakan rumus:

Permutasi: P(n, r) = n! / (n - r)!
Kombinasi: C(n, r) = n! / (r! × (n - r)!)

Hasil perhitungan untuk pasangan (a, c) akan ditampilkan pada baris pertama, sedangkan hasil untuk pasangan (b, d) ditampilkan pada baris kedua.

Sebagai contoh, jika a = 5, c = 3 dan b = 10, d = 10, maka:

P(5,3) = 5!/2! = 120/2 = 60
C(5,3) = 5!/(3! × 2!) = 120/12 = 10
P(10,10) = 10!/0! = 3628800/1 = 3628800
C(10,10) = 10!/(10! × 0!) = 10!/10! = 1]


### 2. [Soal]
#### Soal2.go

```go
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

```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal2.png)
[Program ini digunakan untuk menentukan pemenang dalam kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan.

Program membaca nama peserta satu per satu. Untuk setiap peserta, program akan membaca 8 buah waktu pengerjaan soal (dalam menit). Jika peserta tidak menyelesaikan soal, maka waktu akan bernilai 301 menit.

Perhitungan dilakukan dengan ketentuan:

Jika waktu ≠ 301, maka soal dianggap berhasil diselesaikan dan dihitung sebagai soal yang selesai
Total waktu (skor) adalah jumlah seluruh waktu dari soal yang berhasil diselesaikan

Pemenang ditentukan berdasarkan:

Jumlah soal yang diselesaikan paling banyak
Jika jumlah soal sama, maka dipilih yang memiliki total waktu paling kecil

Hasil akhir yang ditampilkan adalah nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang diperoleh.]


### 3. [Soal]
#### Soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal3.png)
[Program ini meminta pengguna untuk memasukkan sebuah bilangan bulat positif n sebagai nilai awal deret. Setelah itu, program akan membentuk deret bilangan berdasarkan aturan Skiena dan Revilla dengan menggunakan prosedur cetakDeret.

Perhitungan dilakukan dengan ketentuan:

Jika n genap, maka suku berikutnya adalah 1/2n
Jika n ganjil, maka suku berikutnya adalah 3n + 1

Proses ini dilakukan secara berulang hingga mencapai nilai 1 sebagai suku terakhir.]

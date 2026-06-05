# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
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



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da37ae30fd1d8daec240ad079a52a5e794287730/Modul%2012/Output/Soal1.png)
[Program ini dibuat untuk membantu proses penghitungan suara dalam pemilihan ketua RT. Pengguna memasukkan data nomor calon yang dipilih oleh warga secara berurutan hingga ditemukan nilai 0 sebagai penanda akhir data. Setiap data yang masuk dihitung sebagai suara masuk, kemudian divalidasi apakah termasuk nomor calon yang sah, yaitu bernilai antara 1 sampai 20. Suara yang valid akan dicatat dan dihitung jumlahnya menggunakan sebuah array yang menyimpan banyaknya suara untuk masing-masing calon. Setelah seluruh data selesai dibaca, program menampilkan jumlah suara masuk, jumlah suara sah, serta daftar nomor calon yang memperoleh suara beserta banyaknya suara yang diterima oleh setiap calon]


### 2. [Soal]
#### Soal2.go

```go
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

	ketua := 1
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := -1
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if wakil == -1 || suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}

```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal2.1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da37ae30fd1d8daec240ad079a52a5e794287730/Modul%2012/Output/Soal2.1.png)
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soa2.21.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da37ae30fd1d8daec240ad079a52a5e794287730/Modul%2012/Output/Soal2.2.png)
[Program ini dibuat untuk membantu proses rekapitulasi hasil pemilihan ketua RT sekaligus menentukan pasangan ketua dan wakil ketua yang terpilih. Pengguna memasukkan data nomor calon yang dipilih warga hingga ditemukan nilai 0 sebagai penanda akhir masukan. Program akan memvalidasi setiap data sehingga hanya nomor calon antara 1 sampai 20 yang dianggap sebagai suara sah dan dihitung ke dalam array pencatatan suara. Setelah seluruh data diproses, program melakukan penelusuran terhadap jumlah suara setiap calon untuk menentukan calon dengan suara terbanyak sebagai ketua RT. Selanjutnya program mencari calon dengan jumlah suara terbanyak kedua sebagai wakil ketua RT. Sebagai hasil akhir, program menampilkan jumlah suara masuk, jumlah suara sah, nomor calon yang terpilih sebagai ketua RT, dan nomor calon yang terpilih sebagai wakil ketua RT.]


### 3. [Soal]
#### Soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1

	for kr <= kn {
		med := (kr + kn) / 2

		if data[med] == k {
			return med
		} else if data[med] < k {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}

	return -1
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da37ae30fd1d8daec240ad079a52a5e794287730/Modul%2012/Output/Soal3.1.png)
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da37ae30fd1d8daec240ad079a52a5e794287730/Modul%2012/Output/Soal3.2.png)
[Program ini dibuat untuk mencari keberadaan suatu bilangan pada kumpulan data integer yang telah terurut secara membesar (ascending). Pengguna memasukkan jumlah data, nilai yang ingin dicari, serta seluruh elemen array yang sudah dalam keadaan terurut. Data-data tersebut disimpan ke dalam array melalui prosedur pengisian data. Selanjutnya program menggunakan algoritma Binary Search untuk mencari posisi bilangan yang diinginkan. Proses pencarian dilakukan dengan membandingkan nilai yang dicari terhadap elemen tengah array, kemudian mempersempit ruang pencarian ke bagian kiri atau kanan sesuai hasil perbandingan hingga data ditemukan atau ruang pencarian habis. Jika data ditemukan, program menampilkan indeks posisi data tersebut yang dihitung mulai dari 0. Jika data tidak ditemukan, program menampilkan tulisan "TIDAK ADA" sebagai hasil pencarian.]

# <h1 align="center">Laporan Praktikum Modul 7 dan 9 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	center titik
	r      int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.center, p) <= float64(c.r)
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.r)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal1.png)
[Program ini menerima input dua lingkaran (masing-masing berupa koordinat titik pusat dan radius) serta satu titik sembarang. Program menghitung jarak titik tersebut ke masing-masing pusat lingkaran menggunakan rumus jarak.

Jika jarak ≤ radius, maka titik berada di dalam lingkaran. Output program menentukan apakah titik berada di dalam lingkaran 1, lingkaran 2, di dalam keduanya, atau di luar keduanya.]


### 2. [Soal]
#### Soal2.go

```go
package main
import (
	"fmt"
	"math"
)

func main() {
	var n, x, idx, cari int
	var total float64
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Hapus indeks: ")
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Setelah hapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Std Deviasi:", std)

	fmt.Print("Cari angka: ")
	fmt.Scan(&cari)
	count := 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal2.1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal2.1.png))
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soa2.21.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal2.2.png))
[Program ini menerima input sejumlah N bilangan bulat ke dalam sebuah array. Program kemudian menampilkan seluruh isi array, elemen dengan indeks ganjil, dan elemen dengan indeks genap (dengan asumsi indeks ke-0 adalah genap).
Selanjutnya, program menerima nilai x dan menampilkan elemen array pada indeks kelipatan x. Program juga dapat menghapus elemen pada indeks tertentu, lalu menampilkan array hasil penghapusan.
Terakhir, program menghitung dan menampilkan rata-rata, standar deviasi, serta frekuensi kemunculan suatu bilangan tertentu dalam array.]


### 3. [Soal]
#### Soal3.go

```go
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



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal3.png)
[Program ini menerima input nama dua klub sepak bola yang bertanding. Selanjutnya, program menerima input skor dari setiap pertandingan kedua klub secara berulang.

Jika skor salah satu klub lebih besar, maka klub tersebut dinyatakan sebagai pemenang. Jika skor sama, maka hasilnya adalah draw. Proses input berhenti ketika salah satu atau kedua skor bernilai negatif.

Output program menampilkan hasil setiap pertandingan dan diakhiri dengan pesan bahwa pertandingan selesai.]

### 4. [Soal]
#### Soal4.go

```go
package main
import "fmt"

const NMAX = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0

	fmt.Scan(&ch)
	for ch != "." && *n < NMAX {
		t[*n] = rune(ch[0])
		*n++

		fmt.Scan(&ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks: ")
	isiArray(&tab, &n)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal4.1png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal4.1.png)
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal4.2png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/7117f84b95fcf1f32ee02b795b3707425ff1042b/Modul%209/Output/Soal4.2.png)
[Program ini menerima input sejumlah karakter (teks) yang disimpan ke dalam array hingga tanda titik (.) sebagai penanda akhir input. Program kemudian membalik urutan karakter dalam array dan menampilkannya sebagai reverse teks.
Program ini merupakan pengembangan dari versi sebelumnya dengan menambahkan fungsi untuk memeriksa palindrom. Selanjutnya, program mengecek apakah teks tersebut merupakan palindrom, yaitu dibaca sama dari depan maupun belakang. Jika iya, maka ditampilkan "Palindrom", jika tidak maka "Bukan palindrom".]

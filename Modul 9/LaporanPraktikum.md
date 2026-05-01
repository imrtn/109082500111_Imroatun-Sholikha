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
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/90e9f965fcd3cdae379a6da36eb8c832d30f9a43/Modul%203/Output/Soal1.png)
[Program ini meminta input dua lingkaran (titik pusat dan radius) serta satu titik. Program menghitung jarak titik ke masing-masing pusat lingkaran menggunakan rumus jarak.

Jika jarak ≤ radius, maka titik berada di dalam lingkaran.
Hasilnya berupa posisi titik:

di dalam kedua lingkaran
di dalam salah satu
atau di luar keduanya]


### 2. [Soal]
#### Soal2.go

```go
package main
import (
	"fmt"
	"math"
)

func main() {
	var n int
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

	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
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

	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Std Deviasi:", std)

	var cari int
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
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/27f91b3e378a90f936a35e09ffb982bae8ae5401/Modul%203/Output/Soal2.png))
[Program ini mengisi array dengan N bilangan dari pengguna, lalu menampilkan beberapa operasi seperti:

isi array
indeks ganjil & genap
indeks kelipatan x
hapus elemen tertentu
rata-rata
standar deviasi
frekuensi angka tertentu]


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
	pemenang := []string{}
	i := 1

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			fmt.Println("Hasil:", klubA)
			pemenang = append(pemenang, klubA)
		} else if b > a {
			fmt.Println("Hasil:", klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Println("Hasil: Draw")
		}
		i++
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println("Pemenang:", pemenang)
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/27f91b3e378a90f936a35e09ffb982bae8ae5401/Modul%203/Output/Soal3.png)
[Program ini mencatat hasil pertandingan dua klub.
Pengguna memasukkan skor tiap pertandingan.

Program menentukan pemenang:

skor lebih besar = menang
sama = draw

Nama klub pemenang disimpan dalam array.
Input berhenti jika ada skor negatif, lalu ditampilkan daftar pemenang]

### 4. [Soal]
#### Soal4.go

```go
package main
import "fmt"

const NMAX = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c ", &ch)
		if ch == '.' {
			break
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
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

	fmt.Print("Reverse: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/27f91b3e378a90f936a35e09ffb982bae8ae5401/Modul%203/Output/Soal3.png)
[Program ini membaca karakter ke dalam array sampai tanda titik (.) dimasukkan.

Program kemudian:

membalik urutan array
mengecek apakah termasuk palindrom

Palindrom adalah teks yang dibaca sama dari depan maupun belakang]

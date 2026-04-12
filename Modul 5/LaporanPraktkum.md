# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
package main

import "fmt"

func fibonachi(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonachi(n-1) + fibonachi(n-2)
	}
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(fibonachi(i), " ")
	}
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal1.png)

[Program mengimplementasikan deret Fibonacci menggunakan fungsi rekursif.

Deret dimulai dari suku ke-0 = 0 dan ke-1 = 1, sedangkan suku berikutnya merupakan penjumlahan dua suku sebelumnya (Sn = Sn-1 + Sn-2).

Program akan menampilkan deret Fibonacci hingga suku ke-10.]


### 2. [Soal]
#### Soal2.go

```go
package main

import "fmt"

func bintang(n int) {
	if n == 1 {
		fmt.Println("*")
	} else {
		bintang(n - 1)
	}

	for i := 0; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal2.png)
[Program ini digunakan untuk menampilkan pola bintang berbentuk segitiga dengan jumlah baris sesuai input n menggunakan fungsi rekursif.

Setiap baris menampilkan jumlah bintang yang meningkat dari 1 hingga n. Proses dilakukan dengan memanggil fungsi secara berulang hingga mencapai kondisi dasar, kemudian mencetak bintang secara bertahap.

Input berupa satu bilangan bulat n, dan output berupa pola bintang dari 1 sampai n]


### 3. [Soal]
#### Soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	faktor(n, 1)
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal3.png)
[Program ini digunakan untuk menampilkan faktor dari suatu bilangan N dengan menggunakan rekursif.

Program akan mengecek satu per satu bilangan dari 1 sampai N. Jika bilangan tersebut bisa membagi N, maka akan ditampilkan sebagai faktor. Proses ini dilakukan terus secara rekursif sampai semua bilangan diperiksa.

Input berupa satu bilangan bulat positif N, dan output berupa faktor-faktor dari N yang ditampilkan dari yang kecil ke besar]


### 4. [Soal]
#### Soal4.go

```go
package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	ganjil(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	ganjil(n, 1)
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal3.png)
[Program ini digunakan untuk menampilkan barisan bilangan dari N turun hingga 1 kemudian naik kembali ke N menggunakan metode rekursif.
Fungsi rekursif akan mencetak nilai N terlebih dahulu, kemudian memanggil dirinya dengan nilai N-1 hingga mencapai kondisi dasar (N = 1). Setelah itu, program akan mencetak kembali nilai N saat proses kembali, sehingga terbentuk pola turun lalu naik]



### 5. [Soal]
#### Soal5.go

```go
package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	ganjil(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	ganjil(n, 1)
}



```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal3.png)
[Program ini digunakan untuk menampilkan bilangan ganjil dari 1 hingga N menggunakan metode rekursif.
Fungsi rekursif akan mencetak nilai awal 1, kemudian memanggil dirinya dengan menambahkan 2 pada setiap langkah (i + 2), sehingga hanya bilangan ganjil yang ditampilkan hingga mencapai batas N.]



### 6. [Soal]
#### Soal6.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(pangkat(x, y))
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul3/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/a014d1c054cff52d7698e2befa2c539659fb3583/Modul%204/Output/Soal3.png)
[Program ini digunakan untuk menghitung hasil perpangkatan dari dua bilangan bulat x^y menggunakan metode rekursif.
Fungsi rekursif bekerja dengan mengalikan nilai x secara berulang hingga jumlah pangkat (y) mencapai 0. Kondisi dasar terjadi saat y = 0, di mana hasilnya adalah 1. Nilai hasil kemudian dikembalikan melalui proses rekursif hingga diperoleh hasil akhir.]
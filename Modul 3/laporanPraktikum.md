# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
package main
import "fmt"

func factorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

func permutation(n, r int) int {
    return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
    return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(permutation(a, c), combination(a, c))
    fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program ini meminta pengguna memasukkan empat bilangan bulat yaitu a, b, c, dan d. Setelah itu, program menghitung nilai permutasi dan kombinasi dari pasangan (a, c) dan (b, d) menggunakan fungsi faktorial.

Hasil perhitungan untuk (a, c) ditampilkan pada baris pertama, sedangkan hasil untuk (b, d) ditampilkan pada baris kedua. Contohnya, jika a=5, c=3 dan b=10, d=10, maka:
P(5,3) = 5!/2! = 120/2 = 60
C(5,3) = 5!/(3!×2!) = 120/12 = 10
P(10,10) = 10!/0! = 3628800/1 = 3628800
C(10,10) = 10!/(10!×0!) = 1
Sehingga output yang dihasilkan adalah:
60 10
3628800 1]


### 2. [Soal]
#### Soal2.go

```go
package main
import "fmt"

func f(x int) int {
    return x * x
}

func g(x int) int {
    return x - 2
}

func h(x int) int {
    return x + 1
}

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    fmt.Println(f(g(h(a))))
    fmt.Println(g(h(f(b))))
    fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program ini meminta pengguna memasukkan tiga bilangan bulat yaitu a, b, dan c. Program memiliki tiga fungsi yaitu f(x) = x², g(x) = x - 2, dan h(x) = x + 1 yang digunakan untuk melakukan operasi fungsi komposisi.

Setelah input diterima, program menghitung tiga komposisi fungsi yaitu f(g(h(a))), g(h(f(b))), dan h(f(g(c))). Contohnya, jika a=7, b=2, c=10, maka:
(fogoh)(7) = 36
(gohof)(2) = 3
(hofog)(10) = 65
Hasil tersebut kemudian ditampilkan sebagai tiga baris output sesuai urutan perhitungannya..]


### 3. [Soal]
#### Soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	dx := a - c
	dy := b - d
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalamLingkaran1 := didalam(cx1, cy1, r1, x, y)
	dalamLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program ini meminta pengguna memasukkan data dua lingkaran berupa titik pusat dan jari-jari (cx1, cy1, r1 dan cx2, cy2, r2), serta satu titik (x, y). Program menggunakan fungsi jarak untuk menghitung jarak antara dua titik, dan fungsi didalam untuk mengecek apakah titik berada di dalam suatu lingkaran.

Setelah itu, program menentukan posisi titik terhadap kedua lingkaran dan menampilkan hasilnya. Contohnya:
Input: 1 1 5 dan titik 2 2 → Titik di dalam lingkaran 1
Input: 2 1 2 dan 3 4 5 6, titik 7 8 → Titik di dalam lingkaran 2
Input: 3 5 10 dan -15 4 20, titik 0 0 → Titik di dalam lingkaran 1 dan 2
Input: 4 1 1 dan 5 8 8 4, titik 15 20 → Titik di luar lingkaran 1 dan 2]


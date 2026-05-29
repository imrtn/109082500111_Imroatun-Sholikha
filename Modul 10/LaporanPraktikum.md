# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64
	var min, max float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		if i == 0 {
			min = berat[i]
			max = berat[i]
		} else {
			if berat[i] < min {
				min = berat[i]
			}
			if berat[i] > max {
				max = berat[i]
			}
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da275900b51a9368bfdeb77f1f579f1b2eabf047/Modul%2010/Output/Soal1.png)
[Program ini dibuat untuk membantu pendataan berat anak kelinci yang akan dijual ke pasar. Pengguna diminta memasukkan jumlah anak kelinci lalu mencatat berat setiap anak dalam satuan kilogram; semua nilai berat disimpan dalam sebuah array bertipe real. Setelah seluruh data terkumpul, program melakukan penelusuran pada array untuk menentukan berat terkecil dan terbesar. Proses pencarian dimulai dengan menginisialisasi nilai minimum dan maksimum dari elemen pertama, kemudian program membandingkan setiap elemen berikutnya dengan nilai sementara tersebut dan memperbaruinya bila ditemukan nilai yang lebih kecil atau lebih besar. Setelah seluruh elemen diproses, program menampilkan berat anak kelinci yang paling ringan dan paling berat sebagai hasil akhir.]


### 2. [Soal]
#### Soal2.go

```go
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


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal2.1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da275900b51a9368bfdeb77f1f579f1b2eabf047/Modul%2010/Output/Soal2.1.png)
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soa2.21.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da275900b51a9368bfdeb77f1f579f1b2eabf047/Modul%2010/Output/Soal2.2.png)
[Program ini digunakan untuk mengelompokkan data berat ikan ke dalam beberapa wadah berdasarkan kapasitas yang ditetapkan pengguna. Pertama, program menerima masukan berupa jumlah ikan yang akan dijual, kapasitas tiap wadah, dan berat masing-masing ikan; seluruh berat dicatat secara berurutan ke dalam sebuah array. Selanjutnya program membagi data ikan ke wadah secara berurutan sesuai kapasitas yang diberikan, menghitung total berat di setiap wadah dengan menjumlahkan berat ikan yang berada pada wadah tersebut, dan menampilkan hasil penjumlahan untuk tiap wadah kepada pengguna. Setelah itu program menghitung nilai rata-rata berat antar-wadah dengan cara membagi total seluruh berat wadah dengan jumlah wadah yang terbentuk, sehingga pengguna memperoleh gambaran distribusi berat ikan pada setiap wadah beserta nilai rata-ratanya]


### 3. [Soal]
#### Soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata = rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}


```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da275900b51a9368bfdeb77f1f579f1b2eabf047/Modul%2010/Output/Soal3.1.png)
![https://github.com/imrtn/109082500111/blob/main/Modul9/Output/Soal3.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/da275900b51a9368bfdeb77f1f579f1b2eabf047/Modul%2010/Output/Soal3.2.png)
[Program ini dirancang untuk mendukung pencatatan dan analisis data berat badan balita pada kegiatan posyandu. Pengguna memasukkan sejumlah data berat balita yang kemudian disimpan dalam sebuah array, sementara struktur program dirapikan menggunakan subprogram atau fungsi agar proses perhitungan lebih terorganisir dan mudah dipahami. Satu fungsi bertanggung jawab mencari nilai ekstrem—berat minimum dan maksimum—dengan menelusuri array, sedangkan fungsi lain menghitung nilai rata-rata berat seluruh balita berdasarkan jumlah data yang dimasukkan. Setelah seluruh proses selesai, program menampilkan berat balita terkecil, terbesar, dan nilai rata-ratanya dalam satuan kilogram sebagai ringkasan hasil pengolahan data.]

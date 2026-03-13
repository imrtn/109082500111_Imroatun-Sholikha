# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Imroatun Sholikha] - [109082500111]</p>

## Unguided 

### 1. [Soal]
#### Soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![https://github.com/imrtn/109082500111/blob/main/Modul1/Output/Soal1.png](https://github.com/imrtn/109082500111_Imroatun-Sholikha/blob/6c043bac314779e81f09b802761a490c42743331/Modul%201/Output/Soal1.png)
[Program ini meminta pengguna memasukkan tiga buah string yang disimpan dalam variabel satu, dua, dan tiga. Setelah semua input dimasukkan, program menampilkan urutan awal dari ketiga string tersebut sebagai output awal.

Selanjutnya program menukar posisi nilai ketiga variabel dengan bantuan variabel sementara temp. Nilai satu dipindahkan ke temp, kemudian dua menjadi satu, tiga menjadi dua, dan nilai awal satu (yang ada di temp) menjadi tiga. Hasil akhirnya ditampilkan sebagai output akhir, sehingga urutan string berubah dari urutan awalnya.]


### 2. [Soal]
#### Soal2.go

```go
package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Println("Percobaan : ", i)
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)
		if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu" {
			berhasil = false
		}
	}
	fmt.Println("Berhasil : ", berhasil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program ini meminta pengguna memasukkan empat buah string yang disimpan dalam variabel gelas1, gelas2, gelas3, dan gelas4. Program menggunakan perulangan for untuk melakukan proses input sebanyak 5 kali percobaan. Pada setiap percobaan, pengguna diminta memasukkan empat kata yang mewakili warna gelas.

Setelah pengguna memasukkan data, program akan memeriksa apakah urutan warna yang dimasukkan sesuai dengan yang ditentukan, yaitu merah, kuning, hijau, dan ungu. Jika ada satu saja warna yang tidak sesuai dengan urutan tersebut, maka variabel berhasil akan diubah menjadi false.

Setelah semua percobaan selesai dilakukan, program akan menampilkan nilai dari variabel berhasil. Nilai ini menunjukkan apakah semua input yang diberikan selama lima percobaan sesuai dengan urutan warna yang diharapkan atau tidak. Jika semua benar maka hasilnya true, tetapi jika ada yang salah maka hasilnya false.]


### 3. [Soal]
#### Soal3.go

```go
package main

import "fmt"

func main() {
	var totalGram, kg, sisaGram, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&totalGram)
	kg = totalGram / 1000
	sisaGram = totalGram % 1000

	biayaKg = kg * 10000
	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}
	totalBiaya = biayaKg + biayaSisa
	fmt.Printf("Detail berat: % d kg + % d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. % d + Rp. % d\n", biayaKg, biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program ini meminta pengguna memasukkan berat parsel dalam satuan gram yang disimpan pada variabel totalGram. Selanjutnya program menghitung berapa kilogram (kg) dan sisa gram (sisaGram) dari berat tersebut dengan menggunakan operasi pembagian dan modulus.

Setelah itu program menghitung biaya pengiriman. Biaya untuk setiap kilogram adalah Rp10.000, sehingga nilai tersebut dikalikan dengan jumlah kilogram untuk mendapatkan biayaKg. Untuk sisa gram, jika berat parsel lebih dari 10 kg maka sisa gram tidak dikenakan biaya. Namun jika beratnya 10 kg atau kurang, maka sisa gram akan dihitung biayanya, yaitu Rp5 per gram jika sisa gram ≥ 500, dan Rp15 per gram jika sisa gram < 500.

Terakhir, program menjumlahkan biayaKg dan biayaSisa untuk mendapatkan totalBiaya. Program kemudian menampilkan detail berat dalam bentuk kilogram dan gram, detail biaya masing-masing bagian, serta total biaya yang harus dibayar untuk pengiriman parsel tersebut.]



